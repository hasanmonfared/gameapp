package main

import (
	"context"
	"fmt"
	"gameapp/adapter/redis"
	"gameapp/config"
	"gameapp/entity"
	"gameapp/pkg/protobufencoder"
)

func main() {
	cfg := config.Load("config.yml")
	redisAdapter := redis.New(cfg.Redis)
	topic := entity.MatchingUserMatchedEvent
	subscriber := redisAdapter.Client().Subscribe(context.Background(), string(topic))
	for {
		msg, err := subscriber.ReceiveMessage(context.Background())
		if err != nil {
			panic(err)
		}

		switch entity.Event(msg.Channel) {
		case topic:
			processUsersMatchedEvent(msg.Channel, msg.Payload)
		default:
			fmt.Println("invalid topic", msg.Channel)
		}

	}
}
func processUsersMatchedEvent(topic string, data string) {
	mu := protobufencoder.DecodeEventEventMatchingUserMatchedEvent(data)
	fmt.Println("matched users", mu)
	fmt.Println("Received message from " + topic + " topic.")
}
