package delivery

import (
	"errors"
	"net/http"
	"projectly-server/pkg/apierror"
	"projectly-server/internal/domain/user/delivery/token"
	"projectly-server/internal/domain/user/entity"

	"github.com/labstack/echo/v4"
)

type requestUpdate struct {
	Name     string  `json:"name"`
	Surname  string  `json:"surname"`
	Email    string  `json:"email"`
	Language *string `json:"language"`
	Birthday *string `json:"birthday"`
	Location *string `json:"location"`
	About    *string `json:"about"`
}

// Update handles user profile updates.
// @Summary Update user
// @Description	Update user
// @ID				user-update
// @Tags			user
// @Accept			application/json
// @Produce			application/json
// @Param			requestUpdate	body		requestUpdate	true	"Update user details"
// @Success		200				{object}	entity.User		"Updated user"
// @Failure		400				{object}	echo.HTTPError	"Bad request"
// @Failure		500				{object}	echo.HTTPError	"Internal server error"
// @Router			/user/update [post].
func (h *UserHandler) Update(c echo.Context) error {
	var request requestUpdate
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

	targetUser, err := h.UserUseCase.GetUserByEmail(c.Request().Context(), request.Email)
	if err != nil && !errors.Is(err, entity.ErrNoUserFound) {
		return apierror.Validation("Failed to get users")
	}
	if targetUser.ID != 0 && user.Email != request.Email {
		return apierror.Conflict("A user with this email already exists")
	}

	meta := user.Meta
	if request.Language != nil {
		meta.Language = *request.Language
	}
	if request.Birthday != nil {
		meta.Birthday = *request.Birthday
	}
	if request.Location != nil {
		meta.Location = *request.Location
	}
	if request.About != nil {
		meta.About = *request.About
	}

	err = h.UserUseCase.UpdateUser(c.Request().Context(), &entity.User{
		ID:       user.ID,
		Name:     request.Name,
		Surname:  request.Surname,
		Email:    request.Email,
		Password: user.Password,
		Meta:     meta,
	})
	if err != nil {
		return apierror.Internal("Failed to update user")
	}
	return c.NoContent(http.StatusNoContent)
}
