package main

import (
	"github.com/Setsu548/trival-assassement/controllers"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	// routes.InitRoutes(e)
	userController := controllers.NewUserController()
	e.GET("/users", userController.GetUser)
	e.Start(":8000")

}
