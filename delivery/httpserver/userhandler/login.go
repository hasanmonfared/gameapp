package userhandler

import (
	"gameapp/param"
	"gameapp/pkg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) userLogin(c echo.Context) error {

	var req param.LoginRequest
	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if filedErrors, err := h.userValidator.ValidateLoginRequest(req); err != nil {
		msg, code := httpmsg.Error(err)
		return c.JSON(code, echo.Map{
			"message": msg,
			"errors":  filedErrors,
		})
		return echo.NewHTTPError(code, msg, filedErrors)
	}
	resp, err := h.userSvc.Login(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
