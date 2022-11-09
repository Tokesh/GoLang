package utils

import (
	json2 "encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"projectGoLang/source/domain/entity"
	"time"
)

type RedisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, expires time.Duration) *RedisCache {
	return &RedisCache{
		host:    host,
		db:      db,
		expires: expires,
	}
}
func (cache *RedisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{Addr: cache.host, Password: "", DB: cache.db})
}
func (cache *RedisCache) Set(key string, value *entity.Product) {
	client := cache.getClient()

	json, err := json2.Marshal(value)
	if err != nil {
		fmt.Errorf("Set problem", err)
	}
	client.Set(key, json, cache.expires*time.Second)
}

func (cache *RedisCache) Get(key string) *entity.Product {
	client := cache.getClient()
	val, err := client.Get(key).Result()
	fmt.Println(val)
	if err != nil {
		fmt.Errorf("Get problem", err)
	}
	prod := entity.Product{}
	err = json2.Unmarshal([]byte(val), &prod)
	if err != nil {
		fmt.Errorf("Get problem, unmarshalling", err)
	}
	return &prod
}
