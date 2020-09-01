package http

import (
	"UserDataTestTask/models"
	"UserDataTestTask/users"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
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

	_, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	_, err = strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	myUsers, err := h.useCase.GetUsers(c)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	return c.JSON(http.StatusOK, myUsers)
}

func (h *Handler) AddUserHandler(c echo.Context) error {
	userBSON := models.User{}

	// Decode request
	err := json.NewDecoder(c.Request().Body).Decode(&userBSON)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	// close BODY req
	defer func() {
		err = c.Request().Body.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	user, err := h.useCase.AddUser(c, &userBSON)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	return c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUserHandler(c echo.Context) error {
	pushUser := models.User{}

	// decode request
	err := json.NewDecoder(c.Request().Body).Decode(&pushUser)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	// close BODY req
	defer func() {
		err = c.Request().Body.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	user, err := h.useCase.UpdateUser(c, &pushUser)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	return c.JSON(http.StatusOK, user)
}
