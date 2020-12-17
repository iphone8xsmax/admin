package cache

import "github.com/garyburd/redigo/redis"

var RDS RedisDataStore

func Init() *redis.Pool {
	RDS := RedisDataStore{
		RedisHost: "127.0.0.1:6379",
		RedisPwd:  "1q2w3e4r",
		RedisDB:   "1",
		Timeout:   20,
		RedisPool: nil,

	}
	return RDS.NewPool()
}