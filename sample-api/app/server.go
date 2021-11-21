package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"sample-api/app/controller"
	"sample-api/app/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {

	if err := run(); err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {

		return c.String(http.StatusOK, "Hello, world!")
	})

	e.GET("/name/:name", getName)
	e.GET("/show", show)
	e.POST("/save", save)
	e.POST("/saveMultipart", saveMultipart)
	e.POST("/users", getJson)

	e.Logger.Fatal(e.Start(":8080"))
}

// Path Prameters
func getName(c echo.Context) error {

	name := c.Param("name")
	return c.String(http.StatusOK, name)
}

// Query Parameters
func show(c echo.Context) error {

	team := c.QueryParam("team")
	member := c.QueryParam("member")

	s := fmt.Sprintf("team: %s, member: %s", team, member)
	return c.String(http.StatusOK, s)
}

// Form application/x-www-form-urlencoded
func save(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	s := fmt.Sprintf("name: %s, email: %s", name, email)
	return c.String(http.StatusOK, s)
}

// Form multipart/form-data
func saveMultipart(c echo.Context) error {
	name := c.FormValue("name")

	text, err := c.FormFile("text")
	if err != nil {
		return err
	}

	src, err := text.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(text.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.String(http.StatusOK, "Hi, "+name)
}

func getJson(c echo.Context) error {
	// u := new(User)
	u := new(model.User)
	// Content-Typeヘッダーに基づいて, バインドされる
	// ex. applicatin/jsonの場合、{"name": "hoge", "email": "huga"}のデータがバインドされる
	if err := c.Bind(u); err != nil {
		return err
	}
	user := model.User{Name: u.Name, Email: u.Email}
	return c.JSON(http.StatusCreated, user)

}

func run() error {
	log.Println("run begin.")

	// create a user
	// user1 := User{Name: "wdk", Email: "hoge@hogehuga.com"}
	_, err := controller.CreateUser()
	if err != nil {
		log.Println("create user error.")
		return err
	}

	const id = 1
	user, err := controller.GetUser(id)
	if err != nil {
		log.Println("get user error.")
		return err
	}

	// list all users
	users, err := controller.GetUsers()
	if err != nil {
		log.Println(err)
		log.Println("list error.")
		return err
	}

	if err := controller.DeleteUser(id); err != nil {
		log.Println(err)
		log.Println("user delete error.")
		return err
	}
	// Delete After list all users
	usersAfterDeleted, err := controller.GetUsers()
	if err != nil {
		log.Println(err)
		log.Println("list error.")
		return err
	}

	log.Println("user = ", user)
	log.Println("users = ", users)
	log.Println("usersAfterDeleted = ", usersAfterDeleted)
	return nil
}
