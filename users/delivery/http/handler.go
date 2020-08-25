package http

import (
	"UserDataTestTask/users"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
)

type Handler struct {
	useCase users.UseCase
}

func NewHandler(useCase users.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) GetUsersHandler(c echo.Context) error {
	myUsers, err := h.useCase.GetUsers(c)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	return c.JSON(http.StatusOK, myUsers)
}

func (h *Handler) AddUserHandler(c echo.Context) error {
	user, err := h.useCase.AddUser(c)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	return c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUserHandler(c echo.Context) error {
	user, err := h.useCase.UpdateUser(c)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	return c.JSON(http.StatusOK, user)
}
