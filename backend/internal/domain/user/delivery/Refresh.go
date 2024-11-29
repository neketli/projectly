package delivery

import (
	"fmt"
	"net/http"
	"task-tracker-server/internal/domain/user/entity"

	"github.com/labstack/echo/v4"
)

type requestRefresh struct {
	RefreshToken string `json:"refreshToken"`
}

// @Summary		Refresh user token
// @Description	Refresh user token and return new access and refresh tokens
// @Tags			user
// @Accept			application/json
// @Produce		application/json
// @Param			refreshToken	body		requestRefresh	true	"Refresh token value"
// @Success		200				{object}	entity.Tokens
// @Failure		400				{object}	echo.HTTPError
// @Failure		401				{object}	echo.HTTPError
// @Failure		500				{object}	echo.HTTPError
// @Router			/user/refresh [post]
func (h *UserHandler) Refresh(c echo.Context) error {
	var request requestRefresh
	if err := c.Bind(&request); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("validation error: %s", err.Error()),
		}
	}

	user, err := h.UserUseCase.GetUserByRefreshToken(c.Request().Context(), request.RefreshToken)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: fmt.Sprintf("can't find user with this refresh token: %s", err.Error()),
		}
	}

	accessToken, err := h.UserUseCase.CreateAccess(&user)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "can't create access token",
		}
	}
	refreshToken, err := h.UserUseCase.CreateRefresh(&user)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "can't create refresh token",
		}
	}
	return c.JSON(http.StatusOK, &entity.Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	})
}
