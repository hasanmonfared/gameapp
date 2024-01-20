package main

import (
	"context"
	"encoding/base64"
	"gameapp/adapter/redis"
	"gameapp/config"
	"gameapp/contract/golang/matching"
	"gameapp/entity"
	"gameapp/pkg/slice"
	"google.golang.org/protobuf/proto"
)

func main() {
	cfg := config.Load("config.yml")
	redisAdapter := redis.New(cfg.Redis)
	topic := "matching.user_matched"

	mu := entity.MatchUsers{
		Category: entity.FootballCategory,
		UserIDs:  []uint{1, 4},
	}
	pbMu := matching.MatchUsers{
		Category: string(mu.Category),
		UserIds:  slice.MapFromUintToUint64(mu.UserIDs),
	}
	payload, err := proto.Marshal(&pbMu)
	if err != nil {
		panic(err)
	}
	payloadStr := base64.StdEncoding.EncodeToString(payload)
	if err := redisAdapter.Client().Publish(context.Background(), topic, payloadStr).Err(); err != nil {
		panic(err)
	}
}
