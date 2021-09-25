package db

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var (
	rdb *redis.Client
)

func NewRedisClient() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:7002",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	fmt.Println("redis connection successfully created")
	pong, err := rdb.Ping().Result()
	fmt.Println(pong, err)
}

func Set(key string, value interface{}, expires time.Duration) error {
	p, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = rdb.Set(key, string(p), expires).Err()
	if err != nil {
		return err
	}

	return nil
}

func Get(key string, dest interface{}) error {
	val, err := rdb.Get(key).Bytes()
	fmt.Println(key)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(val, dest); err != nil {
		return err
	}

	return nil
}

func Delete(key string) error {
	err := rdb.Del(key).Err()
	if err != nil {
		return err
	}

	return nil
}
