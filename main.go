package main

import (
	"gameapp/adapter/redis"
	"gameapp/config"
	"gameapp/delivery/httpserver"
	"gameapp/repository/migrator"
	"gameapp/repository/mysql"
	"gameapp/repository/mysql/mysqlaccesscontrol"
	"gameapp/repository/mysql/mysqluser"
	"gameapp/repository/redis/redismatching"
	"gameapp/service/authorizationservice"
	"gameapp/service/authservice"
	"gameapp/service/backofficeuserservice"
	"gameapp/service/matchingservice"
	"gameapp/service/userservice"
	"gameapp/validator/matchingvalidator"
	"gameapp/validator/uservalidator"
)

func main() {
	cfg := config.Load("config.yml")
	mgr := migrator.New(cfg.Mysql)
	mgr.Up()
	authSvc, userSvc, userValidator, backofficeSvc, authorizationSvc, matchingSvc, matchingV := setupServices(cfg)
	server := httpserver.New(cfg, authSvc, userSvc, userValidator, backofficeSvc, authorizationSvc, matchingSvc, matchingV)
	server.Serve()
}

func setupServices(cfg config.Config) (authservice.Service, userservice.Service, uservalidator.Validator,
	backofficeuserservice.Service, authorizationservice.Service,
	matchingservice.Service,
	matchingvalidator.Validator) {
	authSvc := authservice.New(cfg.Auth)
	MysqlRepo := mysql.New(cfg.Mysql)

	backofficeUserSvc := backofficeuserservice.New()

	aclMysql := mysqlaccesscontrol.New(MysqlRepo)
	authorizationSvc := authorizationservice.New(aclMysql)

	userMysql := mysqluser.New(MysqlRepo)
	userSvc := userservice.New(authSvc, userMysql)

	redisAdapter := redis.New(cfg.Redis)
	matchingRepo := redismatching.New(redisAdapter)
	matchingSvc := matchingservice.New(cfg.MatchingService, matchingRepo)
	matchingV := matchingvalidator.New()

	uV := uservalidator.New(userMysql)
	return authSvc, userSvc, uV, backofficeUserSvc, authorizationSvc, matchingSvc, matchingV
}
