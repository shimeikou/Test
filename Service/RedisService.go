package Service

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

//const REDIS_IP_PORT = "redis-no-cluster.eu24cn.ng.0001.apne1.cache.amazonaws.com:6379"
const REDIS_IP_PORT = "127.0.0.1:6379"

var RedisPool *redis.Pool

func InitRedisPool() {
	RedisPool = &redis.Pool{
		MaxIdle:     3,
		MaxActive:   6,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", REDIS_IP_PORT) },
	}
}
