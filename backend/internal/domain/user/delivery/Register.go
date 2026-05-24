package delivery

import (
	"errors"
	"net/http"
	"projectly-server/internal/domain/user/entity"
	"projectly-server/pkg/apierror"

	"github.com/labstack/echo/v4"
)

type requestRegister struct {
	Name     string `json:"name" validate:"required,min=2,max=128"`
	Surname  string `json:"surname" validate:"required,min=2,max=128"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Register handles user registration.
// @Summary Register user
// @ID			user-register
// @Tags		auth
// @Accept		application/json
// @Produce		application/json
// @Param		requestRegister	body		requestRegister	true	"register user details"
// @Success	201				{object}	entity.User		"Created user"
// @Failure	400				{object}	echo.HTTPError	"Bad request"
// @Failure	500				{object}	echo.HTTPError	"Internal server error"
// @Router		/auth/register [post].
func (h *UserHandler) Register(c echo.Context) error {
	var request requestRegister
	if err := c.Bind(&request); err != nil {
		return apierror.Validation("Invalid request body")
	}

	if err := c.Validate(request); err != nil {
		return err
	}

	user, err := h.UserUseCase.GetUserByEmail(c.Request().Context(), request.Email)
	if err != nil && !errors.Is(err, entity.ErrNoUserFound) {
		return apierror.Internal("Failed to check user existence")
	}
	if user.ID != 0 {
		return apierror.Conflict("A user with this email already exists")
	}

	err = h.UserUseCase.CreateUser(c.Request().Context(), &entity.User{
		ID:       0,
		Name:     request.Name,
		Surname:  request.Surname,
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return apierror.Internal("Failed to register user")
	}
	return c.NoContent(http.StatusCreated)
}
