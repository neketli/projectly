package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @Summary		Upload task attachments
// @Description	Upload task files and save it in object storage
// @ID				task-create-attachment
// @Tags			task
// @Accept			application/json
// @Produce			application/json
// @Param			files	formData	file	true	"files"
// @Success		200 {object}	[]string 		"File names"
// @Failure		400	{object}	echo.HTTPError	"Bad request"
// @Failure		500	{object}	echo.HTTPError	"Internal server error"
// @Router			/task/{id}/create-attachments [post]
func (h *TaskHandler) CreateAttachment(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("file form validation error: %s", err.Error()),
		}
	}
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid task id",
		}
	}
	files := form.File["files"]
	filenames := make([]string, len(files))

	for index, file := range files {
		if file.Size > 30*1024*1024 {
			return &echo.HTTPError{
				Code:    http.StatusRequestEntityTooLarge,
				Message: fmt.Sprintf("file %s size exceeds 30MB", file.Filename),
			}
		}

		filename, err := h.taskUseCase.CreateAttachment(c.Request().Context(), taskID, file)
		if err != nil {
			return &echo.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: fmt.Sprintf("update task error: %s", err.Error()),
			}
		}

		filenames[index] = filename
	}

	return c.JSON(http.StatusOK, filenames)
}
