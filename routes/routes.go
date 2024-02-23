package routes

import (
	"net/http"

	"github.com/Setsu548/trival-assassement/controllers"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	api := e.Group("/api")
	usersInit(api)

	userController := controllers.NewUserController()
	api.GET("/users", userController.GetUser)
}

func usersInit(api *echo.Group) {
	api.GET("/docs", func(c echo.Context) error { return c.String(http.StatusOK, "ok") })
}
