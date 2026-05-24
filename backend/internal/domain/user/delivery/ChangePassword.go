package delivery

import (
	"errors"
	"net/http"
	"projectly-server/pkg/apierror"
	"projectly-server/internal/domain/user/delivery/token"
	"projectly-server/internal/domain/user/entity"

	"github.com/labstack/echo/v4"
)

type requestChangePassword struct {
	Password string `json:"password"`
}

// ChangePassword handles password change requests.
// @Summary Change user password
// @Description	Change user password by old password
// @ID				user-change-password
// @Tags			user
// @Accept			application/json
// @Produce		application/json
// @Param			requestChangePassword	body	requestChangePassword	true	"Change password request body"
// @Success		200
// @Failure		400	{object}	echo.HTTPError
// @Failure		401	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/user/change-password [post].
func (h *UserHandler) ChangePassword(c echo.Context) error {
	var request requestChangePassword
	if err := c.Bind(&request); err != nil {
		return apierror.Validation("Invalid request body")
	}

	claims, err := token.GetUserClaims(c)
	if err != nil {
		return apierror.Validation("Failed to authenticate user")
	}

	user, err := h.UserUseCase.GetUserByEmail(c.Request().Context(), claims.Email)
	if err != nil && !errors.Is(err, entity.ErrNoUserFound) {
		return apierror.Internal("Failed to get users")
	}
	if user.ID == 0 {
		return apierror.NotFound("User not found")
	}

	err = h.UserUseCase.ChangePassword(c.Request().Context(), &entity.User{
		ID:       user.ID,
		Name:     user.Name,
		Surname:  user.Surname,
		Email:    user.Email,
		Password: request.Password,
	})
	if err != nil {
		return apierror.Internal("Failed to change password")
	}
	return c.NoContent(http.StatusNoContent)
}
