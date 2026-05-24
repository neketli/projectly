package delivery

import (
	"net/http"
	"strconv"

	"projectly-server/pkg/apierror"
	"github.com/labstack/echo/v4"
)

// CreateAttachment handles uploading task attachments.
// @Summary Upload task attachments
// @Description	Upload task files and save it in object storage
// @ID				task-create-attachment
// @Tags			task
// @Accept			application/json
// @Produce			application/json
// @Param			files	formData	file	true	"files"
// @Success		200 {object}	[]string 		"File names"
// @Failure		400	{object}	echo.HTTPError	"Bad request"
// @Failure		500	{object}	echo.HTTPError	"Internal server error"
// @Router			/task/{id}/create-attachments [post].
func (h *TaskHandler) CreateAttachment(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return apierror.Validation("Invalid file form")
	}
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apierror.Validation("Invalid task ID")
	}
	files := form.File["files"]
	filenames := make([]string, len(files))

	for index, file := range files {
		if file.Size > 30*1024*1024 {
			return apierror.Validation("File size exceeds 30MB")
		}

		filename, err := h.taskUseCase.CreateAttachment(c.Request().Context(), taskID, file)
		if err != nil {
			return apierror.Internal("Failed to update task")
		}

		filenames[index] = filename
	}

	return c.JSON(http.StatusOK, filenames)
}
