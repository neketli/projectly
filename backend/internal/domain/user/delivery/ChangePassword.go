package delivery

import (
	"errors"
	"fmt"
	"net/http"
	"task-tracker-server/internal/domain/user/delivery/utils"
	"task-tracker-server/internal/domain/user/entity"

	"github.com/labstack/echo/v4"
)

type requestChangePassword struct {
	Password string `json:"password"`
}

// @Summary     Change user password
// @Description Change user password by old password
// @ID          user-change-password
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       requestChangePassword body requestChangePassword true "Change password request body"
// @Success     200
// @Failure     400  {object} echo.HTTPError
// @Failure     401  {object} echo.HTTPError
// @Failure     500  {object} echo.HTTPError
// @Router      /user/change-password [post]
func (h *UserHandler) ChangePassword(c echo.Context) error {
	var request requestChangePassword
	if err := c.Bind(&request); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("validation error: %s", err.Error()),
		}
	}

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

	err = h.UserUsecase.ChangePassword(c.Request().Context(), &entity.User{
		ID:       user.ID,
		Name:     user.Name,
		Surname:  user.Surname,
		Email:    user.Email,
		Password: request.Password,
	})
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("change password error: %s", err.Error()),
		}
	}
	return c.NoContent(http.StatusNoContent)
}
