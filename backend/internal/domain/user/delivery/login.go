package delivery

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type requestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @Summary     Login user
// @ID          user-login
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Param		requestLogin	body	requestLogin	true	"Login request body"
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
	return c.JSON(http.StatusOK, map[string]string{
		"access":  accessToken,
		"refresh": refreshToken,
	})
}
