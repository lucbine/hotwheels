package core

import (
	"fmt"
	"hotwheels/agent/internal/config"
	"log"
	"time"

	"github.com/go-redis/redis"
)

const errConfigKeyNil = "redis config nil: %s"

var LocalCache *redis.Client

func InitRedis() (err error) {
	LocalCache, err = NewRedisClient("redis.local_cache")
	if err != nil {
		fmt.Printf("redis local_cache init failed" + err.Error())
		return err
	}
	fmt.Println("redis init sucessful")
	return nil
}

func NewRedisClient(key string, ping ...bool) (client *redis.Client, err error) {
	if !config.IsSet(key) {
		return nil, fmt.Errorf(errConfigKeyNil, key)
	}

	client = redis.NewClient(&redis.Options{
		Addr:         config.GetString(key + ".addr"),
		Password:     config.GetString(key + ".password"),
		DB:           config.GetInt(key + ".db"),
		DialTimeout:  config.GetDuration(key+".dialTimeout") * time.Millisecond,
		ReadTimeout:  config.GetDuration(key+".readTimeout") * time.Millisecond,
		WriteTimeout: config.GetDuration(key+".writeTimeout") * time.Millisecond,
		IdleTimeout:  config.GetDuration(key+".idleTimeout") * time.Millisecond,
		MaxRetries:   config.GetInt(key + ".maxRetries"),
		PoolSize:     config.GetInt(key + ".poolSize"),
		MinIdleConns: config.GetInt(key + ".minIdleConns"),
	})

	if len(ping) > 0 && ping[0] {
		if err := client.Ping().Err(); err != nil {
			log.Printf("redis %s ping failed, reply is %s\n", key, err.Error())
		}
	}
	return client, nil
}
