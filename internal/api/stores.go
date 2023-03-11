package api

import (
	"github.com/go-redis/redis/v8"
	"github.com/smallcase/go-be-template/config"
	"github.com/smallcase/go-be-template/pkg/store"
	binottoStore "github.com/smallcase/go-be-template/pkg/store/binotto/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type Stores struct {
	BinottoStore store.BinottoStore
}

func initializeStores(conf *config.Config, mongoDB *mongo.Database, redis *redis.Client) Stores {
	return Stores{
		BinottoStore: binottoStore.New(mongoDB),
	}
}
