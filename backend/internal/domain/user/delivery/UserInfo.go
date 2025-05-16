package delivery

import (
	"errors"
	"net/http"

	"projectly-server/internal/domain/user/entity"

	"github.com/labstack/echo/v4"
)

// @Summary		Get user info by email
// @Description	Get user info by email
// @ID				user-info-by-email
// @Tags			user
// @Accept			application/json
// @Produce			application/json
// @Param			email	path		string	true	"User email"
// @Success		200		{object}	entity.User	"User info"
// @Failure		400		{object}	echo.HTTPError	"Bad request"
// @Failure		404		{object}	echo.HTTPError	"User not found"
// @Failure		500		{object}	echo.HTTPError	"Internal server error"
// @Router			/user/{email} [get]
func (h *UserHandler) UserInfo(c echo.Context) error {
	email := c.Param("email")
	if email == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "email is required",
		}
	}

	user, err := h.UserUseCase.GetUserByEmail(c.Request().Context(), email)
	if err != nil {
		if errors.Is(err, entity.ErrNoUserFound) {
			return &echo.HTTPError{
				Code:    http.StatusNotFound,
				Message: "user not found",
			}
		}

		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return c.JSON(http.StatusOK, user)
}
