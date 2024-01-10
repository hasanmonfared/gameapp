package matchinghandler

import (
	"gameapp/service/authservice"
	"gameapp/service/matchingservice"
	"gameapp/validator/matchingvalidator"
)

type Handler struct {
	authConfig        authservice.Config
	authSvc           authservice.Service
	matchingSvc       matchingservice.Service
	matchingValidator matchingvalidator.Validator
}

func New(authSvc authservice.Service, authConfig authservice.Config, matchingSvc matchingservice.Service, matchingValidator matchingvalidator.Validator) Handler {
	return Handler{
		authConfig:        authConfig,
		authSvc:           authSvc,
		matchingSvc:       matchingSvc,
		matchingValidator: matchingValidator,
	}
}
