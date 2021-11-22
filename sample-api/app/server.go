package main

import (
	"net/http"
	"sample-api/app/controller"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	})

	e.GET("/user/:id", controller.GetUser)
	e.GET("/users", controller.GetUsers)
	e.POST("/user", controller.CreateUser)
	e.PUT("/user/:id", controller.UpdateUser)
	e.DELETE("/user/:id", controller.DeleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}
