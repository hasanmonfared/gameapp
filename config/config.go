package config

import (
	"gameapp/adapter/redis"
	"gameapp/repository/mysql"
	"gameapp/service/authservice"
	"gameapp/service/matchingservice"
	"gameapp/service/presenceservice"
	"time"
)

type Application struct {
	GracefulShutdownTimeout time.Duration `koanf:"graceful_shutdown_timeout"`
}
type HTTPServer struct {
	Port int `koanf:"port"`
}
type Config struct {
	Application     Application            `koanf:"app lication"`
	HTTPServer      HTTPServer             `koanf:"http_server"`
	Auth            authservice.Config     `koand:"auth"`
	Mysql           mysql.Config           `koanf:"mysql"`
	MatchingService matchingservice.Config `koanf:"matching_service"`
	Redis           redis.Config           `koanf:"redis"`
	PresenceService presenceservice.Config `koanf:"presence_service"`
}
