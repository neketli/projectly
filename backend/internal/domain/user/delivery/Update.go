package delivery

import (
	"errors"
	"fmt"
	"net/http"
	"projectly-server/internal/domain/user/delivery/token"
	"projectly-server/internal/domain/user/entity"

	"github.com/labstack/echo/v4"
)

type requestUpdate struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}

// @Summary		Update user
// @Description	Update user
// @ID				user-update
// @Tags			user
// @Accept			application/json
// @Produce		application/json
// @Param			requestUpdate	body		requestUpdate	true	"Update user details"
// @Success		200				{object}	entity.User		"Updated user"
// @Failure		400				{object}	echo.HTTPError	"Bad request"
// @Failure		500				{object}	echo.HTTPError	"Internal server error"
// @Router			/user/update [post]
func (h *UserHandler) Update(c echo.Context) error {
	var request requestUpdate
	if err := c.Bind(&request); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("validation error: %s", err.Error()),
		}
	}

	claims, err := token.GetUserClaims(c)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("can't extract user from token: %s", err.Error()),
		}
	}

	user, err := h.UserUseCase.GetUserByEmail(c.Request().Context(), claims.Email)
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

	targetUser, err := h.UserUseCase.GetUserByEmail(c.Request().Context(), request.Email)
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

	err = h.UserUseCase.UpdateUser(c.Request().Context(), &entity.User{
		ID:       user.ID,
		Name:     request.Name,
		Surname:  request.Surname,
		Email:    request.Email,
		Password: user.Password,
	})
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("update user error: %s", err.Error()),
		}
	}
	return c.NoContent(http.StatusNoContent)
}
