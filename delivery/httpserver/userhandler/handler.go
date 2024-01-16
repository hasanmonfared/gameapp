package userhandler

import (
	"gameapp/service/authservice"
	"gameapp/service/presenceservice"
	"gameapp/service/userservice"
	"gameapp/validator/uservalidator"
)

type Handler struct {
	authConfig    authservice.Config
	authSvc       authservice.Service
	userSvc       userservice.Service
	userValidator uservalidator.Validator
	presenceSVc   presenceservice.Service
}

func New(authSvc authservice.Service, userSvc userservice.Service, userValidator uservalidator.Validator, authConfig authservice.Config, presenceSvc presenceservice.Service) Handler {
	return Handler{
		authConfig:    authConfig,
		authSvc:       authSvc,
		userSvc:       userSvc,
		userValidator: userValidator,
		presenceSVc:   presenceSvc,
	}
}
