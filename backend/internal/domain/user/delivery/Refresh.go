package delivery

import (
	"net/http"
	"projectly-server/internal/domain/user/entity"
	"projectly-server/pkg/apierror"

	"github.com/labstack/echo/v4"
)

type requestRefresh struct {
	RefreshToken string `json:"refreshToken"`
}

// Refresh handles token refresh.
// @Summary Refresh user token
// @Description	Refresh user token and return new access and refresh tokens
// @ID				user-refresh
// @Tags			auth
// @Accept			application/json
// @Produce			application/json
// @Param			refreshToken	body		requestRefresh	true	"Refresh token value"
// @Success		200				{object}	entity.Tokens
// @Failure		400				{object}	echo.HTTPError
// @Failure		401				{object}	echo.HTTPError
// @Failure		500				{object}	echo.HTTPError
// @Router			/auth/refresh [post].
func (h *UserHandler) Refresh(c echo.Context) error {
	var request requestRefresh
	if err := c.Bind(&request); err != nil {
		return apierror.Validation("Invalid request body")
	}

	user, err := h.UserUseCase.GetUserByRefreshToken(c.Request().Context(), request.RefreshToken)
	if err != nil {
		return apierror.Unauthorized("Invalid or expired refresh token")
	}

	accessToken, err := h.UserUseCase.CreateAccess(&user)
	if err != nil {
		return apierror.Internal("Failed to generate access token")
	}
	refreshToken, err := h.UserUseCase.CreateRefresh(&user)
	if err != nil {
		return apierror.Internal("Failed to generate refresh token")
	}
	return c.JSON(http.StatusOK, &entity.Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	})
}
