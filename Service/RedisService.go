package service

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

//REDIS_IP_PORT_AWS aws redis ip&port
//const REDIS_IP_PORT_AWS = "redis-no-cluster.eu24cn.ng.0001.apne1.cache.amazonaws.com:6379"

//RedisIPPortLocal local redis ip
const RedisIPPortLocal = "127.0.0.1:6379"

var (
	//RedisConnectionPool redisコネクションプール
	RedisConnectionPool *redis.Pool
)

//RedisInit Redis関連の初期化
func RedisInit() {
	RedisConnectionPool = RedisInitPool()
}

//RedisInitPool initialize a pool for redis connection
func RedisInitPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", RedisIPPortLocal) },
	}
}
