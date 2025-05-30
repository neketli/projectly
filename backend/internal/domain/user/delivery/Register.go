package delivery

import (
	"errors"
	"fmt"
	"net/http"
	"projectly-server/internal/domain/user/entity"

	"github.com/labstack/echo/v4"
)

type requestRegister struct {
	Name     string `json:"name" validate:"required,min=2,max=128"`
	Surname  string `json:"surname" validate:"required,min=2,max=128"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// @Summary		Register user
// @ID			user-register
// @Tags		auth
// @Accept		application/json
// @Produce		application/json
// @Param		requestRegister	body		requestRegister	true	"register user details"
// @Success	201				{object}	entity.User		"Created user"
// @Failure	400				{object}	echo.HTTPError	"Bad request"
// @Failure	500				{object}	echo.HTTPError	"Internal server error"
// @Router		/auth/register [post]
func (h *UserHandler) Register(c echo.Context) error {
	var request requestRegister
	if err := c.Bind(&request); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("validation error: %s", err.Error()),
		}
	}

	if err := c.Validate(request); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("validation error: %s", err.Error()),
		}
	}

	user, err := h.UserUseCase.GetUserByEmail(c.Request().Context(), request.Email)
	if err != nil && !errors.Is(err, entity.ErrNoUserFound) {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't get users: %s", err.Error()),
		}
	}
	if user.ID != 0 {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "user already exists",
		}
	}

	err = h.UserUseCase.CreateUser(c.Request().Context(), &entity.User{
		ID:       0,
		Name:     request.Name,
		Surname:  request.Surname,
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("registration error: %s", err.Error()),
		}
	}
	return c.NoContent(http.StatusCreated)
}
