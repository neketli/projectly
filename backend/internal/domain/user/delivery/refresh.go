package delivery

import (
	"fmt"
	"net/http"
	"task-tracker-server/internal/domain/user/entity"

	"github.com/labstack/echo/v4"
)

type requestRefresh struct {
	RefreshToken string `json:"refreshToken"`
}

// @Summary     Refresh user token
// @ID          user-refresh
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Param		refreshToken	body	requestRefresh	true	"Refresh token value"
// @Router      /user/refresh [post]
func (h *UserHandler) Refresh(c echo.Context) error {
	var request requestRefresh
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("validation error: %s", err.Error()))
	}

	user, err := h.UserUsecase.GetUserByRefreshToken(c.Request().Context(), request.RefreshToken)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, entity.ErrNoUserFound)
	}

	accessToken, err := h.UserUsecase.CreateAccess(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Ошибка при создании JWT токена для пользователя")
	}
	refreshToken, err := h.UserUsecase.CreateRefresh(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Ошибка при создании JWT токена для пользователя")
	}
	return c.JSON(http.StatusOK, map[string]string{
		"access":  accessToken,
		"refresh": refreshToken,
	})
}
