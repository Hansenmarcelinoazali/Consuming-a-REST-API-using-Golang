package repository

import (
	"external_api/redis"
	"fmt"
	"log"
	"time"
)

func SetRedis(url string, productdata string) {

	rdb := redis.NewRedisClient()
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
