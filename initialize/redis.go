package initialize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var (
	RedisCli *redis.Client
)

func redisInit() {
	RedisCli = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
		PoolSize: 100,
	})
	_, err := RedisCli.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln("Redis connect error:", err)
	}
	// 初始化短地址的偏移
	// TODO: key 应该是个全局变量
	_ = RedisCli.SetNX(context.Background(), "next.url.id", 1000000000, 0)
}
