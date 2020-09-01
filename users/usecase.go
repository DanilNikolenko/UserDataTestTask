package users

import (
	"UserDataTestTask/models"
	"github.com/labstack/echo"
)

type UseCase interface {
	GetUsers(c echo.Context) (*[]models.User, error)
	AddUser(c echo.Context, user *models.User) (*models.User, error)
	UpdateUser(c echo.Context, user *models.User) (*models.User, error)
}
