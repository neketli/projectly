package delivery

import (
	"errors"
	"fmt"
	"net/http"
	"task-tracker-server/internal/domain/user/entity"

	"github.com/labstack/echo/v4"
)

type requestRegister struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @Summary     Register user
// @ID          user-register
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Param		requestRegister	body	requestRegister	true	"register user json"
// @Router      /user/register [post]
func (h *UserHandler) Register(c echo.Context) error {
	var request requestRegister
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("validation error: %s", err.Error()))
	}

	user, err := h.UserUsecase.GetUserByEmail(c.Request().Context(), request.Email)
	if err != nil && !errors.Is(err, entity.ErrNoUserFound) {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("can't get users: %s", err.Error()))
	}
	if user.ID != 0 {
		return c.JSON(http.StatusBadRequest, "user already exists")
	}
	err = h.UserUsecase.CreateUser(c.Request().Context(), &entity.User{
		ID:       0,
		Name:     request.Name,
		Surname:  request.Surname,
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("registration error: %s", err.Error()))
	}
	return c.NoContent(http.StatusCreated)
}
