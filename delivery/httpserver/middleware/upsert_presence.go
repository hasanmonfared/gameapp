package middleware

import (
	"gameapp/param"
	"gameapp/pkg/claim"
	"gameapp/pkg/errmsg"
	timestamp "gameapp/pkg/timsestamp"
	"gameapp/service/presenceservice"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UpsertPresence(service presenceservice.Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := claim.GetClaimsFromContext(c)
			_, err := service.Upsert(c.Request().Context(), param.UpsertPresenceRequest{
				UserID:    claims.UserID,
				Timestamp: timestamp.Now(),
			})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": errmsg.ErrorMsgSomethingWentWrong,
				})
			}
			return next(c)
		}

	}
}
