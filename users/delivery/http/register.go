package http

import (
	"UserDataTestTask/users"
	"fmt"
	"github.com/labstack/echo"
)

func RegisterHTTPEndpoints(e *echo.Echo, uc users.UseCase) {
	h := NewHandler(uc)

	e.GET("/getUsers", h.GetUsersHandler)
	e.POST("/addUser", h.AddUserHandler)
	e.POST("/updateUser", h.UpdateUserHandler)

	fmt.Println("Methods registered!")
}
