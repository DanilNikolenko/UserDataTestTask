package usecase

import (
	"UserDataTestTask/models"
	"UserDataTestTask/users"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type UsersUseCase struct {
	UsersStorage users.UsersRepository
}

func NewUsersUseCase(UR users.UsersRepository) *UsersUseCase {
	return &UsersUseCase{
		UsersStorage: UR,
	}
}

func (r *UsersUseCase) GetUsers(c echo.Context) (*[]models.User, error) {
	myUsers, err := r.UsersStorage.GetUsersFromDB(c)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return myUsers, nil
}

func (r *UsersUseCase) AddUser(c echo.Context) (*models.User, error) {
	user, err := r.UsersStorage.AddUserToDB(c)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return user, nil
}

func (r *UsersUseCase) UpdateUser(c echo.Context) (*models.User, error) {
	user, err := r.UsersStorage.UpdateUserInDB(c)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return user, nil
}
