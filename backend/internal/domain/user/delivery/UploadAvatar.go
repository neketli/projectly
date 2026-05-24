package delivery

import (
	"errors"
	"net/http"
	"path/filepath"
	"projectly-server/pkg/apierror"
	"projectly-server/internal/domain/user/delivery/token"
	"projectly-server/internal/domain/user/entity"

	"github.com/labstack/echo/v4"
)

// UploadAvatar handles user avatar upload.
// @Summary Upload user avatar
// @Description	Upload user file and save it in object storage
// @ID				user-upload-avatar
// @Tags			user
// @Accept			application/json
// @Produce			application/json
// @Param			image	formData	file	true	"file"
// @Success		200
// @Failure		400	{object}	echo.HTTPError	"Bad request"
// @Failure		500	{object}	echo.HTTPError	"Internal server error"
// @Router			/user/upload-avatar [post].
func (h *UserHandler) UploadAvatar(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return apierror.Validation("Invalid avatar form")
	}
	file := form.File["image"][0]

	if file.Size > 2*1024*1024 {
		return apierror.Validation("File size exceeds 2MB")
	}

	availableExtensions := map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
	}

	_, foundExtension := availableExtensions[filepath.Ext(file.Filename)]
	if !foundExtension {
		return apierror.Validation("File extension is not supported")
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

	err = h.UserUseCase.UploadAvatar(c.Request().Context(), user, file)
	if err != nil {
		return apierror.Internal("Failed to update user")
	}
	return c.NoContent(http.StatusNoContent)
}
