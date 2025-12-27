package cache

import (
	"context"
	"github.com/hewo/tik-shop/shared/consts"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"sync"
	"time"
)

var (
	GlobalRedisClient *redis.Client
	once              sync.Once
)

func ConnectRedis(path string) (*redis.Client, error) {
	if err := godotenv.Load(path); err != nil {
		return nil, err
	}

	addr := os.Getenv("REDIS_ADDR")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")

	client := redis.NewClient(&redis.Options{
		Addr:         addr + ":" + port,
		Password:     password,
		DB:           0,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		PoolSize:     10,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return client, nil

}

func InitRedis() {
	client, err := ConnectRedis(consts.RPCRedisEnvPath)
	if err != nil {
		panic(err)
	}
	GlobalRedisClient = client
	log.Println("Connected to redis successfully")
}

func GetRedisClient() *redis.Client {
	once.Do(func() {
		InitRedis()
	})
	return GlobalRedisClient
}
