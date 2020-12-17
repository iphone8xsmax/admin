package cache

import (
	"github.com/garyburd/redigo/redis"
	"gowith/logging"
	"time"
)

type RedisDataStore struct {
	RedisHost string
	RedisDB   string
	RedisPwd  string
	Timeout   int64
	RedisPool *redis.Pool
}

//新建连接池
func (r *RedisDataStore) NewPool() *redis.Pool {

	return &redis.Pool{
		Dial:        r.RedisConnect,
		MaxIdle:     10,
		MaxActive:   0,
		IdleTimeout: 1 * time.Second,
		Wait:        true,
	}
}

//建立Redis连接
func (r *RedisDataStore) RedisConnect() (redis.Conn, error) {
	c, err := redis.Dial("tcp", r.RedisHost)
	if err != nil {
		return nil, err
	}
	_, err = c.Do("AUTH", r.RedisPwd)

	if err != nil {
		return nil, err
	}

	_, err = c.Do("SELECT", r.RedisDB)
	if err != nil {
		return nil, err
	}

	redis.DialConnectTimeout(time.Duration(r.Timeout) * time.Second)
	redis.DialReadTimeout(time.Duration(r.Timeout) * time.Second)
	redis.DialWriteTimeout(time.Duration(r.Timeout) * time.Second)

	return c, nil
}


func (r *RedisDataStore) Get(k string) (interface{}, error) {
	c := r.RedisPool.Get()
	defer c.Close()
	v, err := c.Do("GET", k)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (r *RedisDataStore) Set(k, v string) error {
	c := r.RedisPool.Get()
	defer c.Close()
	_, err := c.Do("SET", k, v)
	return err
}

func (r *RedisDataStore) Delete(k string) error {
	c := r.RedisPool.Get()
	defer c.Close()
	_, err := c.Do("DEL", k)
	return err
}

func (r *RedisDataStore) SetEx(k string, v interface{}, ex int64) error {
	c := r.RedisPool.Get()
	defer c.Close()
	_, err := c.Do("SET", k, v, "EX", ex)
	return err
}

func (r *RedisDataStore) IsExist(k string) bool {
	c := r.RedisPool.Get()
	defer c.Close()
	isExist, err := redis.Bool(c.Do("EXISTS", k))
	if err != nil{
		logging.Fatal(err)
	}
	return isExist
}


