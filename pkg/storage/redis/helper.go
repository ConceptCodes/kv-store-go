package redis

import (
	"kv-store/internal/constants"
	"kv-store/pkg/logger"
	"time"
)

var log = logger.GetLogger()

func (r *Redis) GetData(key string) (string, error) {
	log.
		Debug().
		Str("key", key).
		Msgf("Getting key %s from redis", key)

	return r.Client.Get(r.Ctx, key).Result()
}

func (r *Redis) SetData(key string, value string, ttl time.Duration) error {
	log.
		Debug().
		Str("key", key).
		Msgf("Setting key %s in redis", key)

	return r.Client.Set(r.Ctx, key, value, ttl|constants.DefaultRedisTtl).Err()
}

func (r *Redis) Exists(key string) bool {
	log.
		Debug().
		Str("key", key).
		Msgf("Checking if key %s exists in redis", key)

	return r.Client.Exists(r.Ctx, key).Val() == 1
}

func (r *Redis) Expire(key string) error {
	log.
		Debug().
		Str("key", key).
		Msgf("Setting ttl for key %s in redis", key)

	return r.Client.Expire(r.Ctx, key, 0).Err()
}
