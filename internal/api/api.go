package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/smallcase/go-be-template/config"
	"github.com/smallcase/go-be-template/pkg/log"
	"go.mongodb.org/mongo-driver/mongo"
)

func Initialize(conf *config.Config, mongoDB *mongo.Database, redis *redis.Client) {
	if !conf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.Use(gin.Recovery())
	// TODO: middleware to log incoming HTTP request
	// TODO: add CORS and security header middlewares
	stores := initializeStores(conf, mongoDB, redis)
	if err := initializeRouter(router, conf, stores); err != nil {
		log.App.Fatal().Err(err).Msg("Could not initialize Gin router")
	}
}
