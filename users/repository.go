package users

import (
	"UserDataTestTask/models"
	"github.com/labstack/echo"
)

type UsersRepository interface {
	GetUsersFromDB(c echo.Context) (*[]models.User, error)
	AddUserToDB(c echo.Context, user *models.User) (*models.User, error)
	UpdateUserInDB(c echo.Context, user *models.User) (*models.User, error)
}
