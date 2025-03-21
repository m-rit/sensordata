package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

func initdb() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatal("Error connecting to Redis:", err)
		return nil
	}
	return rdb

}
