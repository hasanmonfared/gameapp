package main

import (
	"gameapp/adapter/redis"
	"gameapp/config"
	"gameapp/delivery/grpcserver/presenceserver"
	"gameapp/repository/redis/redispresence"
	"gameapp/service/presenceservice"
)

func main() {
	cfg := config.Load("config.yml")
	redisAdapter := redis.New(cfg.Redis)
	presenceRepo := redispresence.New(redisAdapter)
	presenceSvc := presenceservice.New(cfg.PresenceService, presenceRepo)
	server := presenceserver.New(presenceSvc)
	server.Start()
}
