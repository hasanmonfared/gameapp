package matchinghandler

import (
	"gameapp/param"
	"gameapp/pkg/claim"
	"gameapp/pkg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) addToWaitingList(c echo.Context) error {

	var req param.AddToWaitingListRequest
	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest)
	}
	claims := claim.GetClaimsFromContext(c)
	req.UserID = claims.UserID

	if filedErrors, err := h.matchingValidator.ValidateAddToWaitingList(req); err != nil {
		msg, code := httpmsg.Error(err)
		return c.JSON(code, echo.Map{
			"message": msg,
			"errors":  filedErrors,
		})
		return echo.NewHTTPError(code, msg, filedErrors)
	}

	resp, err := h.matchingSvc.AddToWaitingList(req)
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}
	return c.JSON(http.StatusOK, resp)
}
