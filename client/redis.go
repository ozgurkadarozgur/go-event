package client

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
)

type redisClient struct {
	connection *redis.Client
}

func NewRedisClient() Client {
	return &redisClient{
		connection: getConnection(),
	}
}

func getConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func (c *redisClient) Publish(channel string, message []byte) {
	c.connection.Publish(ctx, channel, message)
}

func (c *redisClient) Subscribe(channel string, callback func(payload string)) {
	pubSub := c.connection.Subscribe(ctx, channel)
	go handlePubSub(pubSub, callback)
}

func handlePubSub(pubSub *redis.PubSub, callback func(payload string)) {
	for message := range pubSub.Channel() {
		callback(message.Payload)
	}
}
