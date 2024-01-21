package main

import (
	"context"
	"gameapp/adapter/redis"
	"gameapp/config"
	"gameapp/entity"
	"gameapp/pkg/protobufencoder"
)

func main() {
	cfg := config.Load("config.yml")
	redisAdapter := redis.New(cfg.Redis)
	topic := entity.MatchingUserMatchedEvent

	mu := entity.MatchUsers{
		Category: entity.FootballCategory,
		UserIDs:  []uint{1, 4},
	}
	payload := protobufencoder.EncodeEvent(topic, mu)
	if err := redisAdapter.Client().Publish(context.Background(), string(topic), payload).Err(); err != nil {
		panic(err)
	}
}
