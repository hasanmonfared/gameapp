package main

import (
	"gameapp/config"
	"gameapp/delivery/httpserver"
	"gameapp/repository/mysql"
	"gameapp/service/authservice"
	"gameapp/service/userservice"
	"time"
)

const (
	JwtSignKey                 = "hgfhhkgg844hf"
	AccessTokenSubject         = "ac"
	RefreshTokenSubject        = "rt"
	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7
)

func main() {
	cfg := config.Config{
		HTTPServer: config.HTTPServer{Port: 8787},
		Auth: authservice.Config{
			SignKey:               JwtSignKey,
			AccessExpirationTime:  AccessTokenExpireDuration,
			RefreshExpirationTime: RefreshTokenExpireDuration,
			AccessSubject:         AccessTokenSubject,
			RefreshSubject:        RefreshTokenSubject,
		},
		Mysql: mysql.Config{
			Username: "gameapp",
			Password: "wontbeused",
			Port:     3320,
			Host:     "127.0.0.1",
			DBName:   "gameapp_db",
		},
	}
	authSvc, userSvc := setupServices(cfg)
	server := httpserver.New(cfg, authSvc, userSvc)
	server.Serve()
}

func setupServices(cfg config.Config) (authservice.Service, userservice.Service) {
	authSvc := authservice.New(cfg.Auth)
	MysqlRepo := mysql.New(cfg.Mysql)
	userSvc := userservice.New(authSvc, MysqlRepo)
	return authSvc, userSvc
}
