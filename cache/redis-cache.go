package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"manager/configs"
	"manager/utils"
	"time"

	"github.com/go-redis/redis/v8"
)

// collection of fields
type redisCache struct {
	host    string
	db      int           //index between 0 and 15
	expires time.Duration // expiration time for all elements in cache in seconds
}

// init creates a new instance of redisCache
func NewRedisCache(host string, db int, exp time.Duration) FlagCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

// create a new redis client
func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", configs.GetEnvVar("REDIS_HOST"), configs.GetEnvVar("REDIS_PORT")),
		Password: configs.GetEnvVar("REDIS_PW"),
		DB:       cache.db,
	})
}

// Implementation of Set - associate flag json to the key
func (cache *redisCache) Set(key string, value interface{}) {
	client := cache.getClient()

	// serialize the flag
	json, err := json.Marshal(value)
	if err != nil {
		utils.HandleErr(err, "Set from redis: marshalling error")
		return
	}

	// set the key to marshalled data
	err = client.Set(context.TODO(), key, json, cache.expires*time.Second).Err()
	if err != nil {
		utils.HandleErr(err, "Error writing to redis cache...")
	}
}

// asynchronously flush all keys from cache
func (cache *redisCache) FlushAllAsync() {
	client := cache.getClient()

	_, err := client.Ping(context.TODO()).Result()
	if err != nil {
		utils.HandleErr(err, "error with cache.getClient() when flushing cache")
		return
	}

	client.FlushAllAsync(context.TODO())
}
