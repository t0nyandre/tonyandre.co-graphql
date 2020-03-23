package redis

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
)

func OpenRedis() (*redis.Client, error) {
	store := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s", os.Getenv("REDIS_URL")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err := store.Ping().Result()
	if err != nil {
		log.Printf("Retry redis connection in 5 seconds ...")
		time.Sleep(time.Duration(5) * time.Second)
		return OpenRedis()
	}

	log.Println("Connection to redis was a success!")

	return store, nil
}
