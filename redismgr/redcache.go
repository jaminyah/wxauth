package redismgr

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// Store email pass in redis
func StoreEmailPass(addr, passwd string) {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Set(ctx, addr, passwd, 2*time.Hour).Result()
	if err != nil {
		panic(err)
	}
}

// Fetch email pass from redis
func FetchPass(addr string) string {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	addr = addr + "-pwd"

	val, err := client.Get(ctx, addr).Result()
	if err != nil {
		fmt.Println("redis get error: ", err)
		val = "expired"
	}
	fmt.Println("email-pwd", val)
	return val
}

// Store email activation code in redis
func StoreEmailCode(addr, code string) {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Set(ctx, addr, code, 30*time.Minute).Result()
	if err != nil {
		panic(err)
	}
}

// ValidateCode compares user auth code to cached
func ValidateCode(addr, userCode string) bool {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	cachedCode, err := client.Get(ctx, addr).Result()
	if err != nil {
		fmt.Println("redis get error: ", err)
	}
	fmt.Printf("email: %s, code: %s", addr, cachedCode)

	var success bool = false
	if userCode == cachedCode {
		success = true
	}
	return success
}
