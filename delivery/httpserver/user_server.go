package httpserver

import (
	"gameapp/service/userservice"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s Server) userLogin(c echo.Context) error {

	var req userservice.LoginRequest
	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := s.userSvc.Login(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
func (s Server) userRegister(c echo.Context) error {
	var req userservice.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := s.userSvc.Register(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, resp)
}
func (s Server) userProfile(c echo.Context) error {
	c.Request().Header.Get("Authorization")
	authToekn := c.Request().Header.Get("Authorization")

	claims, err := s.authSvc.ParseToken(authToekn)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	resp, err := s.userSvc.Profile(userservice.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}
