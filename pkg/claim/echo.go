package claim

import (
	"gameapp/config"
	"gameapp/service/authservice"
	"github.com/labstack/echo/v4"
)

func GetClaimsFromContext(c echo.Context) *authservice.Claims {
	return c.Get(config.AuthMiddlewareContextKey).(*authservice.Claims)
}
