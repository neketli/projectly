package delivery

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type requestRefresh struct {
	RefreshToken string `json:"refreshToken"`
}

// @Summary     Refresh user token
// @ID          user-refresh
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Param		refreshToken	body	requestRefresh	true	"Refresh token value"
// @Router      /user/refresh [post]
func (h *UserHandler) Refresh(c echo.Context) error {
	var request requestRefresh
	if err := c.Bind(&request); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("validation error: %s", err.Error()),
		}
	}

	user, err := h.UserUsecase.GetUserByRefreshToken(c.Request().Context(), request.RefreshToken)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: fmt.Sprintf("can't find user with this refresh token: %s", err.Error()),
		}
	}

	accessToken, err := h.UserUsecase.CreateAccess(&user)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "can't create access token",
		}
	}
	refreshToken, err := h.UserUsecase.CreateRefresh(&user)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "can't create refresh token",
		}
	}
	return c.JSON(http.StatusOK, map[string]string{
		"access":  accessToken,
		"refresh": refreshToken,
	})
}
