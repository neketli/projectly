package delivery

import (
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
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user, err := h.UserUsecase.GetUserByEmail(c.Request().Context(), request.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// Сравнение хэша и пароля
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		return c.JSON(http.StatusUnauthorized, "incorrect email or password")
	}
	accessToken, err := h.UserUsecase.CreateAccess(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "login error")
	}
	refreshToken, err := h.UserUsecase.CreateRefresh(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "login error")
	}
	return c.JSON(http.StatusOK, map[string]string{
		"access":  accessToken,
		"refresh": refreshToken,
	})
}
