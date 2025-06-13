package delivery

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary		Get media
// @ID			media-get
// @Tags		media
// @Produce 	application/octet-stream
// @Param 		filepath path string true "Filepath"
// @Header  200	{string} 	ETag "File ETag"
// @Success 200 {file} file "File content" Headers(ETag=string,Content-Type=string,Content-Disposition=string)
// @Failure	400	{object}	echo.HTTPError	"Invalid filepath"
// @Failure	500	{object}	echo.HTTPError	"Internal server error"
// @Router		/media/{filepath} [get]
func (h *MediaHandler) GetMedia(c echo.Context) error {
	filepath := c.Param("*")
	if filepath == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid filepath",
		}
	}

	file, err := h.mediaUseCase.GetFile(c.Request().Context(), filepath)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't get media: %s", err.Error()),
		}
	}

	c.Response().Header().Set("Content-Disposition", "attachment; filename="+file.FileInfo.Name)
	c.Response().Header().Set("Content-Type", file.FileInfo.ContentType)
	c.Response().Header().Set("ETag", file.FileInfo.ETag)

	return c.Stream(http.StatusOK, file.FileInfo.ContentType, file.Reader)
}
