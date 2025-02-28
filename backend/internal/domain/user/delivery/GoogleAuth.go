package delivery

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary		Google user auth
// @Description	Google user auth login
// @ID				user-google-login
// @Tags			user
// @Accept			application/json
// @Produce			application/json
// @Success		307
// @Failure		400				{object}	echo.HTTPError
// @Failure		500				{object}	echo.HTTPError
// @Router			/auth/google/login [get]
func (h *UserHandler) GoogleLogin(c echo.Context) error {
	redirectURL := h.UserUseCase.GoogleLogin(c.Request().Context(), "/api/v1/auth/google/callback")
	return c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}

// @Summary		Google user auth callback
// @Description	Google user auth callback and redirect with access and refresh tokens
// @ID				user-google-login-callback
// @Tags			user
// @Accept			application/json
// @Produce			application/json
// @Success		307
// @Failure		400				{object}	echo.HTTPError
// @Failure		500				{object}	echo.HTTPError
// @Router			/auth/google/callback [get]
func (h *UserHandler) GoogleCallback(c echo.Context) error {
	code := c.QueryParam("code")
	user, err := h.UserUseCase.GoogleCallback(c.Request().Context(), code)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("can't login: %s", err.Error()),
		}
	}

	accessToken, err := h.UserUseCase.CreateAccess(user)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "login error",
		}
	}
	refreshToken, err := h.UserUseCase.CreateRefresh(user)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "login error",
		}
	}

	redirectUrl := fmt.Sprintf("/auth/login?access=%s&refresh=%s", accessToken, refreshToken)

	return c.Redirect(http.StatusTemporaryRedirect, redirectUrl)
}
