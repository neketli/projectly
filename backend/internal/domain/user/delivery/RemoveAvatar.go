package delivery

import (
	"errors"
	"net/http"
	"projectly-server/pkg/apierror"
	"projectly-server/internal/domain/user/delivery/token"
	"projectly-server/internal/domain/user/entity"

	"github.com/labstack/echo/v4"
)

// RemoveAvatar handles user avatar removal.
// @Summary Remove user avatar
// @Description	Remove user avatar from storage
// @Tags			user
// @Accept			application/json
// @Produce		application/json
// @Success		200
// @Failure		400	{object}	echo.HTTPError	"Bad request"
// @Failure		500	{object}	echo.HTTPError	"Internal server error"
// @Router			/user/remove-avatar [delete].
func (h *UserHandler) RemoveAvatar(c echo.Context) error {
	claims, err := token.GetUserClaims(c)
	if err != nil {
		return apierror.Validation("Failed to authenticate user")
	}

	user, err := h.UserUseCase.GetUserByEmail(c.Request().Context(), claims.Email)
	if err != nil && !errors.Is(err, entity.ErrNoUserFound) {
		return apierror.Internal("Failed to get users")
	}
	if user.ID == 0 {
		return apierror.NotFound("User not found")
	}
	if user.Meta == nil {
		return apierror.NotFound("User has no avatar")
	}

	err = h.UserUseCase.RemoveAvatar(c.Request().Context(), user)
	if err != nil {
		return apierror.Internal("Failed to remove avatar")
	}

	return c.NoContent(http.StatusOK)
}
