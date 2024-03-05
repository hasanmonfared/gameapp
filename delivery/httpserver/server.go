package httpserver

import (
	"fmt"
	"gameapp/config"
	"gameapp/delivery/httpserver/backofficeuserhandler"
	"gameapp/delivery/httpserver/matchinghandler"
	"gameapp/delivery/httpserver/userhandler"
	"gameapp/logger"
	"gameapp/service/authorizationservice"
	"gameapp/service/authservice"
	"gameapp/service/backofficeuserservice"
	"gameapp/service/matchingservice"
	"gameapp/service/presenceservice"
	"gameapp/service/userservice"
	"gameapp/validator/matchingvalidator"
	"gameapp/validator/uservalidator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type Server struct {
	config                config.Config
	userHandler           userhandler.Handler
	backofficeUserHandler backofficeuserhandler.Handler
	matchingHandler       matchinghandler.Handler
	Router                *echo.Echo
}

func New(config config.Config, authSvc authservice.Service, userSvc userservice.Service, userValidator uservalidator.Validator,
	backofficeUserSvc backofficeuserservice.Service, authorizationSvc authorizationservice.Service,
	matchingSvc matchingservice.Service,
	matchingValidator matchingvalidator.Validator,
	presenceSvc presenceservice.Service) Server {
	return Server{
		config:                config,
		userHandler:           userhandler.New(authSvc, userSvc, userValidator, config.Auth, presenceSvc),
		backofficeUserHandler: backofficeuserhandler.New(authSvc, backofficeUserSvc, config.Auth, authorizationSvc),
		matchingHandler:       matchinghandler.New(authSvc, config.Auth, matchingSvc, matchingValidator, presenceSvc),

		Router: echo.New(),
	}
}

func (s Server) Serve() *echo.Echo {
	// Middleware

	s.Router.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:           true,
		LogStatus:        true,
		LogHost:          true,
		LogRequestID:     true,
		LogRemoteIP:      true,
		LogMethod:        true,
		LogContentLength: true,
		LogResponseSize:  true,
		LogLatency:       true,
		LogError:         true,
		LogProtocol:      true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			errMsg := ""
			if v.Error != nil {
				errMsg = v.Error.Error()
			}
			logger.Logger.Named("http-server").Info("request",
				zap.String("request_id", v.RequestID),
				zap.String("host", v.Host),
				zap.String("content_length", v.ContentLength),
				zap.String("protocol", v.Protocol),
				zap.String("method", v.Method),
				zap.Duration("latency", v.Latency),
				zap.String("error", errMsg),
				zap.String("remote_ip", v.RequestID),
				zap.Int64("response_size", v.ResponseSize),
				zap.String("uri", v.URI),
				zap.Int("status", v.Status),
			)
			return nil
		},
	}))

	//s.Router.Use(middleware.Logger())
	s.Router.Use(middleware.Recover())
	s.Router.Use(middleware.RequestID())
	// Routes
	s.Router.GET("/health-check", s.healthCheck)

	s.userHandler.SetUserRoutes(s.Router)
	s.backofficeUserHandler.SetRoutes(s.Router)
	s.matchingHandler.SetRoutes(s.Router)

	address := fmt.Sprintf(":%d", s.config.HTTPServer.Port)
	fmt.Printf("start echo server on %s\n", address)
	if err := s.Router.Start(address); err != nil {
		fmt.Println("router start error", err)
	}
	return s.Router
}
