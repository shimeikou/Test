package RedisUtil

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

const REDIS_IP_PORT = "redis-no-cluster.eu24cn.ng.0001.apne1.cache.amazonaws.com:6379"

var (
	RedisClient redis.Client
)

func Init() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:         REDIS_IP_PORT,
		Password:     "", // no password set
		DB:           0,  // use default DB
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	pong, err := RedisClient.Ping().Result()

	fmt.Println(pong, err)
	if err != nil {
		RedisClient = nil
	}
}

func RedisSet(key string, value interface{}) {
	err := RedisClient.Set(key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func RedisGet(key string) interface{} {
	val, err := RedisClient.Get(key).Result()
	if err == redis.Nil {
		fmt.Println("%S key does not exist", key)
		return nil
	} else if err != nil {
		panic(err)
	} else {
		return val
	}
}
