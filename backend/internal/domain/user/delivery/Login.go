package delivery

import (
	"fmt"
	"net/http"
	"task-tracker-server/internal/domain/user/entity"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type requestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @Summary     Login user
// @Description Login user
// @ID          user-login
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       requestLogin body requestLogin true "Login request body"
// @Success     200          {object} entity.Tokens
// @Failure     400          {object} echo.HTTPError
// @Failure     500          {object} echo.HTTPError
// @Router      /user/login [post]
func (h *UserHandler) Login(c echo.Context) error {
	var request requestLogin
	if err := c.Bind(&request); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("validation error: %s", err.Error()),
		}
	}

	user, err := h.UserUsecase.GetUserByEmail(c.Request().Context(), request.Email)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	// Сравнение хэша и пароля
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		return &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "incorrect email or password",
		}
	}

	accessToken, err := h.UserUsecase.CreateAccess(&user)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "login error",
		}
	}
	refreshToken, err := h.UserUsecase.CreateRefresh(&user)
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
