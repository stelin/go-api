package redis

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func init() {

	fmt.Println("init")

	redisPool := &redis.Pool{
		MaxIdle:   6,
		MaxActive: 26, // max number of connections
		Dial:      dailCallback,
	}
	RedisClient = redisPool
}

func dailCallback() (redis.Conn, error) {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis connect error=", err)
		panic(err.Error())
	}
	return c, err
}

var RedisClient *redis.Pool

func Get(key string) interface{} {
	con := RedisClient.Get()
	defer con.Close()

	result, err := con.Do("GET", key)
	if err != nil {
		fmt.Println("redis get error, error=", err)
		return nil
	}

	fmt.Println("redis get opertation, ret=", result)
	return result
}

func Set(key string, value interface{}, expire int64) error {
	con := RedisClient.Get()
	defer con.Close()

	var err error
	result, err := con.Do("SETEX", key, expire, value)
	if err != nil {
		fmt.Println("redis set is error=", err)
	}

	fmt.Println("redis set ret=", result)
	return err
}
