package connections

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/smallcase/go-be-template/config"
	"github.com/smallcase/go-be-template/pkg/log"
)

func NewRedisClient(ctx context.Context, conf config.RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		DB:   0,
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, errors.Wrap(err, "Could not test Redis connection stability")
	}
	log.Redis.Info().Msg("Successfully established Redis client connection")
	return client, nil
}

func DisconnectRedis(redis *redis.Client) {
	err := redis.Close()
	if err != nil {
		log.Redis.Error().Err(err).Msg("Could not terminate Redis connection :/")
		return
	}
	log.Redis.Info().Msg("Successfully disconnected from Redis")
}
