package main

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"

	sqls "sample-api/app/db"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {

	// if err := run(); err != nil {
	// 	log.Fatal(err)
	// }

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

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

type CreateUser struct {
	Id    int
	Name  string
	Email string
}

func getJson(c echo.Context) error {
	u := new(User)
	// Content-Typeヘッダーに基づいて, バインドされる
	// ex. applicatin/jsonの場合、{"name": "hoge", "email": "huga"}のデータがバインドされる
	if err := c.Bind(u); err != nil {
		return err
	}
	user := CreateUser{Id: 0, Name: u.Name, Email: u.Email}
	return c.JSON(http.StatusCreated, user)

}

func run() error {
	ctx := context.Background()
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		return err
	}

	// list all users
	queries := sqls.New(db)
	users, err := queries.ListUsers(ctx)
	if err != nil {
		return err
	}
	log.Println(users)

	// create a user
	result, err := queries.CreateUser(ctx, sqls.CreateUserParams{
		Name:  "koki wada",
		Email: sql.NullString{String: "xxxx@yyyy.com", Valid: true},
	})
	if err != nil {
		return err
	}

	insertedUserID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	log.Println(insertedUserID)

	// get user we just inserted
	fetchedUser, err := queries.GetUser(ctx, insertedUserID)
	if err != nil {
		return err
	}

	// print true
	log.Println(reflect.DeepEqual(insertedUserID, fetchedUser.ID))
	return nil
}
