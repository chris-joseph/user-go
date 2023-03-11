package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/smallcase/go-be-template/config"
	"github.com/smallcase/go-be-template/internal/api"
	"github.com/smallcase/go-be-template/internal/connections"
	"github.com/smallcase/go-be-template/pkg/log"
	"go.mongodb.org/mongo-driver/mongo"
)

func createNewConnections(conf *config.Config) (*mongo.Database, *redis.Client, error) {
	mongoDB, err := connections.NewMongoClient(context.TODO(), &conf.Mongo)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Could not establish connection with MongoDB")
	}
	redis, err := connections.NewRedisClient(context.TODO(), conf.Redis)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Could not establish connection with Redis")
	}
	return mongoDB, redis, nil
}

func handleAbruptTermination(mongoDB *mongo.Database, redis *redis.Client) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-signals
		log.App.Warn().Msgf("Application abruptly terminating. Received '%s' OS signal.", sig.String())
		wg := sync.WaitGroup{}
		wg.Add(2)
		go func() {
			connections.DisconnectMongoDB(mongoDB)
			wg.Done()
		}()
		go func() {
			connections.DisconnectRedis(redis)
			wg.Done()
		}()
		wg.Wait()
		log.App.Info().Msg("All remote connections disconnected. Killing the application now")
		os.Exit(1)
	}()
}

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	log.Initialize(conf.ApplicationName)
	log.App.Info().Msg("Warming up the tyres. Booting up the application.")
	mongoDB, redis, err := createNewConnections(conf)
	if err != nil {
		log.App.Fatal().Err(err).Msg("Could not establish remote connections")
	}
	log.App.Info().Msg("Tyres are all warmed up. All connections established successfully!")
	handleAbruptTermination(mongoDB, redis)
	api.Initialize(conf, mongoDB, redis)
}
