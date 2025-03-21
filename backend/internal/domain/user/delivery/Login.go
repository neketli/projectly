package delivery

import (
	"fmt"
	"net/http"
	"projectly-server/internal/domain/user/entity"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type requestLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// @Summary		Login user
// @Description	Login user
// @ID				user-login
// @Tags			auth
// @Accept			application/json
// @Produce			application/json
// @Param			requestLogin	body		requestLogin	true	"Login request body"
// @Success		200				{object}	entity.Tokens
// @Failure		400				{object}	echo.HTTPError
// @Failure		500				{object}	echo.HTTPError
// @Router			/auth/login [post]
func (h *UserHandler) Login(c echo.Context) error {
	var request requestLogin
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
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		return &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "incorrect email or password",
		}
	}

	accessToken, err := h.UserUseCase.CreateAccess(&user)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "login error",
		}
	}
	refreshToken, err := h.UserUseCase.CreateRefresh(&user)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "login error",
		}
	}
	return c.JSON(http.StatusOK, &entity.Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	})
}
