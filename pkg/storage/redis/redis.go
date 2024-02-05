package redis

import (
	"context"
	"fmt"
	"sync"

	"kv-store/config"
	"kv-store/internal/constants"
	"kv-store/pkg/logger"

	"github.com/redis/go-redis/v9"
)

var (
	once        sync.Once
	redisClient *redis.Client
)

type Redis struct {
	Client *redis.Client
	Ctx    context.Context
}

func New() *Redis {
	once.Do(func() {
		log := logger.GetLogger()
		log.Debug().Msg("Connecting to redis")

		redisClient = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", config.AppConfig.RedisHost, config.AppConfig.RedisPort),
			Password: config.AppConfig.RedisPassword,
			DB:       config.AppConfig.RedisDB,
		})

		ctx := context.Background()

		_, err := redisClient.Ping(ctx).Result()
		if err != nil {
			log.Error().Err(err).Msg("Error while connecting to redis")
		}
	})

	return &Redis{
		Client: redisClient,
		Ctx:    context.Background(),
	}
}

func (r *Redis) HealthCheck() bool {
	log := logger.GetLogger()
	log.
		Debug().
		Msgf(fmt.Sprintf(constants.HealthCheckMessage, "redis"))

	res := r.Client.Ping(r.Ctx).Err()

	if res != nil {
		log.
			Error().
			Err(res).
			Msgf(constants.HealthCheckError, "redis")
		return false
	}

	log.Info().Msg("Redis is up")

	return true
}

func (r *Redis) Close() error {
	logger.
		GetLogger().
		Debug().
		Msg("Closing redis connection")
	return r.Client.Close()
}
