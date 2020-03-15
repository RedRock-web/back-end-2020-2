package dao

import (
	"gopkg.in/redis.v4"
)

var G_client *redis.Client

func CreateClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()

	return client, err
}
