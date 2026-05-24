package delivery

import (
	"errors"
	"net/http"

	"projectly-server/pkg/apierror"
	"projectly-server/internal/domain/user/entity"

	"github.com/labstack/echo/v4"
)

// UserInfo handles retrieval of user info by email.
// @Summary Get user info by email
// @Description	Get user info by email
// @ID				user-info-by-email
// @Tags			user
// @Accept			application/json
// @Produce			application/json
// @Param			email	path		string	true	"User email"
// @Success		200		{object}	entity.User	"User info"
// @Failure		400		{object}	echo.HTTPError	"Bad request"
// @Failure		404		{object}	echo.HTTPError	"User not found"
// @Failure		500		{object}	echo.HTTPError	"Internal server error"
// @Router			/user/{email} [get].
func (h *UserHandler) UserInfo(c echo.Context) error {
	email := c.Param("email")
	if email == "" {
		return apierror.Validation("Email is required")
	}

	user, err := h.UserUseCase.GetUserByEmail(c.Request().Context(), email)
	if err != nil {
		if errors.Is(err, entity.ErrNoUserFound) {
			return apierror.NotFound("user not found")
		}

	return apierror.Internal("Failed to get user")
	}

	return c.JSON(http.StatusOK, user)
}
