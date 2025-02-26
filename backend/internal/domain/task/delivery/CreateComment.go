package delivery

import (
	"fmt"
	"net/http"
	"projectly-server/internal/domain/user/delivery/token"
	"strconv"

	"github.com/labstack/echo/v4"
)

type createCommentRequest struct {
	Text string `json:"text"`
}

// @Summary		Create comment in task
// @Description	Creates comment in task
// @ID				task-create-comment
// @Tags			task
// @Accept			application/json
// @Produce			application/json
// @Success		201
// @Failure		400	{object}	echo.HTTPError	"Bad request"
// @Failure		500	{object}	echo.HTTPError	"Internal server error"
// @Router			/task/{id}/create-comment [post]
func (h *TaskHandler) CreateComment(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("invalid task id: %s", err.Error()),
		}
	}

	var request createCommentRequest
	if err := c.Bind(&request); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("validation error: %s", err.Error()),
		}
	}

	claims, err := token.GetUserClaims(c)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("can't extract user from token: %s", err.Error()),
		}
	}

	err = h.taskUseCase.CreateComment(c.Request().Context(), taskID, claims.ID, request.Text)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("can't create comment: %s", err.Error()),
		}
	}

	return c.NoContent(http.StatusCreated)
}
