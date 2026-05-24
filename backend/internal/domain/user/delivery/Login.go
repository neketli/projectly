package delivery

import (
	"net/http"
	"projectly-server/internal/domain/user/entity"
	"projectly-server/pkg/apierror"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type requestLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Login handles user authentication.
// @Summary Login user
// @Description	Login user
// @ID				user-login
// @Tags			auth
// @Accept			application/json
// @Produce			application/json
// @Param			requestLogin	body		requestLogin	true	"Login request body"
// @Success		200				{object}	entity.Tokens
// @Failure		400				{object}	echo.HTTPError
// @Failure		500				{object}	echo.HTTPError
// @Router			/auth/login [post].
func (h *UserHandler) Login(c echo.Context) error {
	var request requestLogin
	if err := c.Bind(&request); err != nil {
		return apierror.Validation("Invalid request body")
	}

	if err := c.Validate(request); err != nil {
		return err
	}

	user, err := h.UserUseCase.GetUserByEmail(c.Request().Context(), request.Email)
	if err != nil {
		return apierror.Unauthorized("Incorrect email or password")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		return apierror.Unauthorized("Incorrect email or password")
	}

	accessToken, err := h.UserUseCase.CreateAccess(&user)
	if err != nil {
		return apierror.Internal("Failed to generate access token")
	}
	refreshToken, err := h.UserUseCase.CreateRefresh(&user)
	if err != nil {
		return apierror.Internal("Failed to generate refresh token")
	}
	return c.JSON(http.StatusOK, &entity.Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	})
}
