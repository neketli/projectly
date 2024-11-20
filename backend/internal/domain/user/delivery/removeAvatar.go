package delivery

import (
	"errors"
	"fmt"
	"net/http"
	"task-tracker-server/internal/domain/user/delivery/utils"
	"task-tracker-server/internal/domain/user/entity"

	"github.com/labstack/echo/v4"
)

// @Summary     remove user avatar
// @ID          user-remove-avatar
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Router      /user/remove-avatar [delete]
func (h *UserHandler) RemoveAvatar(c echo.Context) error {
	claims, err := utils.GetUserClaims(c)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("can't extract user from token: %s", err.Error()),
		}
	}

	user, err := h.UserUsecase.GetUserByEmail(c.Request().Context(), claims.Email)
	if err != nil && !errors.Is(err, entity.ErrNoUserFound) {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't get users: %s", err.Error()),
		}
	}
	if user.ID == 0 {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "can't find user",
		}
	}
	if user.Meta == nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "user has no avatar",
		}
	}

	err = h.UserUsecase.RemoveAvatar(c.Request().Context(), user)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("remove user error: %s", err.Error()),
		}
	}

	return c.NoContent(http.StatusOK)
}
