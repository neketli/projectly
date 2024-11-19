package delivery

import (
	"errors"
	"fmt"
	"net/http"
	"task-tracker-server/internal/domain/user/delivery/utils"
	"task-tracker-server/internal/domain/user/entity"

	"github.com/labstack/echo/v4"
)

type requestUpdate struct {
	Name     string  `json:"name"`
	Surname  string  `json:"surname"`
	Email    string  `json:"email"`
	Password *string `json:"password"`
}

// @Summary     Update user
// @ID          user-update
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Param		requestUpdate	body	requestUpdate	true	"update user json"
// @Router      /user/update [post]
func (h *UserHandler) Update(c echo.Context) error {
	var request requestUpdate
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

	targetUser, err := h.UserUsecase.GetUserByEmail(c.Request().Context(), request.Email)
	if err != nil && !errors.Is(err, entity.ErrNoUserFound) {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("can't get users: %s", err.Error()),
		}
	}
	if targetUser.ID != 0 && user.Email != request.Email {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "can't change emails: user with this email already exists",
		}
	}

	password := user.Password
	if request.Password != nil {
		password = *request.Password
	}
	err = h.UserUsecase.UpdateUser(c.Request().Context(), &entity.User{
		ID:       user.ID,
		Name:     request.Name,
		Surname:  request.Surname,
		Email:    request.Email,
		Password: password,
	}, request.Password != nil)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("update user error: %s", err.Error()),
		}
	}
	return c.NoContent(http.StatusNoContent)
}
