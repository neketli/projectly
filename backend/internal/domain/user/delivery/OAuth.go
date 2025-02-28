package delivery

import (
	"fmt"
	"net/http"
	"projectly-server/internal/domain/user/entity"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

// @Summary		Oauth user auth
// @Description	Oauth user auth login
// @ID				user-oauth-login
// @Tags			user
// @Accept			application/json
// @Produce			application/json
// @Success		307
// @Failure		500				{object}	echo.HTTPError
// @Router			/auth/{provider} [get]
func (h *UserHandler) OauthLogin(c echo.Context) error {
	provider := c.Param("provider")
	if provider == "" {
		return c.String(http.StatusBadRequest, "Provider not specified")
	}
	q := c.Request().URL.Query()
	q.Add("provider", c.Param("provider"))
	c.Request().URL.RawQuery = q.Encode()

	gothic.BeginAuthHandler(c.Response(), c.Request())
	return nil
}

// @Summary		Oauth user auth callback
// @Description	Oauth user auth callback and redirect with access and refresh tokens
// @ID				user-oauth-callback
// @Tags			user
// @Accept			application/json
// @Produce			application/json
// @Success		307
// @Failure		400				{object}	echo.HTTPError
// @Failure		500				{object}	echo.HTTPError
// @Router			/auth/{provider}/callback [get]
func (h *UserHandler) OauthCallback(c echo.Context) error {
	gothUser, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("authentication failed: %s", err.Error()),
		}
	}

	user := &entity.User{
		Email:   gothUser.Email,
		Name:    gothUser.FirstName,
		Surname: gothUser.LastName,
		Meta: &entity.UserMeta{
			Provider:   gothUser.Provider,
			ProviderID: gothUser.UserID,
		},
	}

	if user.Name == "" {
		user.Name = gothUser.NickName
	}
	if user.Surname == "" {
		user.Surname = gothUser.NickName
	}

	err = h.UserUseCase.CompleteUserAuth(c.Request().Context(), user)
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

	redirectUrl := fmt.Sprintf("/auth?access=%s&refresh=%s", accessToken, refreshToken)

	return c.Redirect(http.StatusTemporaryRedirect, redirectUrl)
}
