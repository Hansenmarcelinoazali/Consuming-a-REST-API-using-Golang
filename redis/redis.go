package redis

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	// "github.com/go-redis/redis"
)

func NewRedisClient() *redis.Client {
	var redisHost = "localhost:6379"
	var redisPassword = ""
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       0,
	})
	return client
}

func SetRedis(url string, productdata string) {

	rdb := NewRedisClient()
	fmt.Println("redis client initialized")

	key := url
	data := productdata
	ttl := time.Duration(15) * time.Second

	// store data using SET command
	op1 := rdb.Set(key, data, ttl)
	if err := op1.Err(); err != nil {
		// fmt.Printf("unable to SET data. error: %v", err)
		return
	}
	log.Println("set operation success")

}


