package delivery

import (
	"net/http"
	"projectly-server/pkg/apierror"
	"projectly-server/internal/domain/user/delivery/token"
	"strconv"

	"github.com/labstack/echo/v4"
)

type createCommentRequest struct {
	Text string `json:"text"`
}

// CreateComment handles creating a comment in a task.
// @Summary Create comment in task
// @Description	Creates comment in task
// @ID				task-create-comment
// @Tags			task
// @Accept			application/json
// @Produce			application/json
// @Success		201
// @Failure		400	{object}	echo.HTTPError	"Bad request"
// @Failure		500	{object}	echo.HTTPError	"Internal server error"
// @Router			/task/{id}/create-comment [post].
func (h *TaskHandler) CreateComment(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apierror.Validation("Invalid task ID")
	}

	var request createCommentRequest
	if bindErr := c.Bind(&request); bindErr != nil {
		return apierror.Validation("Invalid request body")
	}

	claims, err := token.GetUserClaims(c)
	if err != nil {
		return apierror.Validation("Failed to authenticate user")
	}

	err = h.taskUseCase.CreateComment(c.Request().Context(), taskID, claims.ID, request.Text)
	if err != nil {
		return apierror.Internal("Failed to create comment: %s")
	}

	return c.NoContent(http.StatusCreated)
}
