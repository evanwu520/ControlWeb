package redis

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func NewRedis() *newReis {
	return &newReis{}
}

type newReis struct {
}

func InitRedis() {

	if client != nil {
		client.Close()
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:         "35.221.206.171:56379",
		Password:     "",               // 没有密码，默认值
		DB:           10,               // 默认DB 0
		PoolSize:     10,               // number of connections in the pool
		MinIdleConns: 5,                // minimum number of idle connections
		PoolTimeout:  30 * time.Second, // time to wait for connection if all busy
	})

	// Ping to ensure connection works
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}

	client = rdb

	// rdb.Get("GM:10002:1")
}

func (*newReis) GetClient() *redis.Client {

	return client
}
