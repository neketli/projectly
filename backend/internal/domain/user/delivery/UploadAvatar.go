package delivery

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"task-tracker-server/internal/domain/user/delivery/token"
	"task-tracker-server/internal/domain/user/entity"

	"github.com/labstack/echo/v4"
)

// @Summary     Upload user avatar
// @Description Upload user file and save it in object storage
// @ID          user-upload-avatar
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Param		image formData file true "file"
// @Success     200
// @Failure     400 {object} echo.HTTPError "Bad request"
// @Failure     500 {object} echo.HTTPError "Internal server error"
// @Router      /user/upload-avatar [post]
func (h *UserHandler) UploadAvatar(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("avatar form validation error: %s", err.Error()),
		}
	}
	file := form.File["image"][0]

	if file.Size > 2*1024*1024 {
		return &echo.HTTPError{
			Code:    http.StatusRequestEntityTooLarge,
			Message: "file size exceeds 2MB",
		}
	}

	availableExtensions := map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
	}

	_, foundExtension := availableExtensions[filepath.Ext(file.Filename)]
	if !foundExtension {
		return &echo.HTTPError{
			Code:    http.StatusUnsupportedMediaType,
			Message: "not available extension",
		}
	}

	claims, err := token.GetUserClaims(c)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("can't extract user from token: %s", err.Error()),
		}
	}

	user, err := h.UserUseCase.GetUserByEmail(c.Request().Context(), claims.Email)
	if err != nil && !errors.Is(err, entity.ErrNoUserFound) {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't get users: %s", err.Error()),
		}
	}
	if user.ID == 0 {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "can't find user",
		}
	}

	err = h.UserUseCase.UploadAvatar(c.Request().Context(), user, file)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("update user error: %s", err.Error()),
		}
	}
	return c.NoContent(http.StatusNoContent)
}
