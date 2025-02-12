package database

import (
	// "fmt"
	// "os"

	"os"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func NewRedisClient() {

	// opt, err := redis.ParseURL(os.Getenv("Redis"))
	// if err != nil {
	// 	panic(fmt.Errorf("Fatal error connect Redis: %w \n", err))
	// }
	// fmt.Println("Redis connect success")

	// client := redis.NewClient(opt)
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       0,                           // use default DB
	})

	RDB = client
}
