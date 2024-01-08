package backofficeuserhandler

import (
	"gameapp/delivery/httpserver/middleware"
	"gameapp/entity"
	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {

	userGroup := e.Group("/backoffice/users")

	userGroup.GET("/profile", h.ListUsers, middleware.Auth(h.authSvc, h.authConfig),
		middleware.AccessCheck(h.authorizationSvc, entity.UserListPermission))
}
