package main

import (
	"gameapp/config"
	"gameapp/delivery/httpserver"
	"gameapp/repository/migrator"
	"gameapp/repository/mysql"
	"gameapp/service/authservice"
	"gameapp/service/userservice"
	"gameapp/validator/uservalidator"
)

func main() {
	//cfg2 := config.Load("config.yml")
	//fmt.Println("cfg2", cfg2)
	cfg := config.Config{
		HTTPServer: config.HTTPServer{Port: 8787},
		Auth: authservice.Config{
			SignKey:               config.JwtSignKey,
			AccessExpirationTime:  config.AccessTokenExpireDuration,
			RefreshExpirationTime: config.RefreshTokenExpireDuration,
			AccessSubject:         config.AccessTokenSubject,
			RefreshSubject:        config.RefreshTokenSubject,
		},
		Mysql: mysql.Config{
			Username: "gameapp",
			Password: "wontbeused",
			Port:     3320,
			Host:     "127.0.0.1",
			DBName:   "gameapp_db",
		},
	}
	mgr := migrator.New(cfg.Mysql)
	mgr.Up()
	authSvc, userSvc, userValidator := setupServices(cfg)
	server := httpserver.New(cfg, authSvc, userSvc, userValidator)
	server.Serve()
}

func setupServices(cfg config.Config) (authservice.Service, userservice.Service, uservalidator.Validator) {
	authSvc := authservice.New(cfg.Auth)
	MysqlRepo := mysql.New(cfg.Mysql)
	userSvc := userservice.New(authSvc, MysqlRepo)
	uV := uservalidator.New(MysqlRepo)
	return authSvc, userSvc, uV
}
