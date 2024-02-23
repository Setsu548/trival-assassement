package controllers

import (
	"net/http"

	"github.com/Setsu548/trival-assassement/models"
	"github.com/Setsu548/trival-assassement/services"
	"github.com/labstack/echo/v4"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (cc UserController) GetUser(c echo.Context) error {
	user := new(models.User)

	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "error binding user")
	}

	service, _ := services.GetUsers()

	return c.JSON(http.StatusOK, service)

}
