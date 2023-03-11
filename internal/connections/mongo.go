package connections

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/smallcase/go-be-template/config"
	"github.com/smallcase/go-be-template/pkg/log"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoLog = log.Mongo

func getMonitor() *event.CommandMonitor {
	return &event.CommandMonitor{
		Succeeded: func(ctx context.Context, cse *event.CommandSucceededEvent) {
			MongoLog.Info().Fields(map[string]interface{}{
				"connectionId": cse.ConnectionID,
				"requestId":    cse.RequestID,
				"serviceId":    cse.ServiceID,
				"command":      cse.CommandName,
				"duration":     cse.DurationNanos,
			}).Msg("Mongo query succeeded")
		},
		Failed: func(ctx context.Context, cfe *event.CommandFailedEvent) {
			MongoLog.Error().Fields(map[string]interface{}{
				"connectionId": cfe.ConnectionID,
				"requestId":    cfe.RequestID,
				"serviceId":    cfe.ServiceID,
				"command":      cfe.CommandName,
				"duration":     cfe.DurationNanos,
				"failure":      cfe.Failure,
			}).Msg("Mongo query failed")
		},
	}
}

func NewMongoClient(ctx context.Context, config *config.MongoConfig) (*mongo.Database, error) {
	monitor := getMonitor()
	opts := options.Client().SetMonitor(monitor).ApplyURI(config.HostURI)
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctxWithTimeout, opts) // timeout connection in 10 seconds if not successful
	if err != nil {
		return nil, errors.Wrap(err, "Could not connect with MongoDB client")
	}
	MongoLog.Info().Msg("Pinging MongoDB to verify connection stability")
	if err := client.Ping(ctx, nil); err != nil {
		return nil, errors.Wrap(err, "Could not test MongoDB connection stability")
	}
	MongoLog.Info().Msg("Successfully established MongoDB client connection")
	db := client.Database(config.DBName)
	return db, nil
}

func DisconnectMongoDB(mongoDB *mongo.Database) {
	err := mongoDB.Client().Disconnect(context.TODO())
	if err != nil {
		log.Mongo.Error().Err(err).Msg("Could not terminate MongoDB connection :/")
		return
	}
	log.Mongo.Info().Msg("Successfully disconnected from MongoDB")
}
