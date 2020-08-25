package server

import (
	"UserDataTestTask/services"
	"UserDataTestTask/users"
	"UserDataTestTask/users/delivery/http"
	"UserDataTestTask/users/repository/mongodb"
	"UserDataTestTask/users/usecase"
	"fmt"
	"github.com/labstack/echo"
)

type App struct {
	httpServer *echo.Echo

	AppUC users.UseCase
}

func NewApp() *App {
	conn := services.ConnToMongo()
	repo := mongodb.NewMongoRepository(conn, "usersdb", "users")

	return &App{
		AppUC: usecase.NewUsersUseCase(repo),
	}
}

func (a *App) Run(port string) error {

	e := echo.New()

	http.RegisterHTTPEndpoints(e, a.AppUC)

	fmt.Println("starting server at " + port)

	e.Logger.Fatal(e.Start(port))

	return nil
}
