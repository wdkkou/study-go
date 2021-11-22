package controller

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	sqls "sample-api/app/db"
	"sample-api/app/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	ctx := context.Background()
	db, err := connect(ctx)
	if err != nil {
		log.Println("connect failed.")
		log.Fatal(err.Error())
	}
	defer db.Close()
	queries := sqls.New(db)
	u := new(model.User)
	// Content-Typeヘッダーに基づいて, バインドされる
	// ex. applicatin/jsonの場合、{"name": "hoge", "email": "huga"}のデータがバインドされる
	if err := c.Bind(u); err != nil {
		return err
	}
	user := model.User{Name: u.Name, Email: u.Email}
	params := sqls.CreateUserParams{
		Name:  u.Name,
		Email: u.Email,
	}
	_, err = queries.CreateUser(ctx, params)
	if err != nil {
		log.Println("create user error.")
		return err
	}

	return c.String(http.StatusOK, fmt.Sprintf("created user =%v", user))
}

func GetUser(c echo.Context) error {
	log.Println("GetUser begin.")
	ctx := context.Background()
	db, err := connect(ctx)
	if err != nil {
		log.Println("connect failed.")
		log.Fatal(err.Error())
	}
	defer db.Close()

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	queries := sqls.New(db)
	user, err := queries.GetUser(ctx, id)
	if err != nil {
		log.Println("get user error.")
		return c.String(http.StatusNotFound, fmt.Sprintf("not found id = %d", id))
	}

	str := fmt.Sprintf("name = %v, email = %v\n", user.Name, user.Email)
	log.Println(str)
	return c.String(http.StatusOK, str)
}

func GetUsers(c echo.Context) error {
	ctx := context.Background()
	db, err := connect(ctx)
	if err != nil {
		log.Println("connect failed.")
		log.Fatal(err.Error())
	}
	defer db.Close()
	queries := sqls.New(db)

	users, err := queries.ListUsers(ctx)
	if err != nil {
		log.Println("get users error.")
		return err
	}

	str := fmt.Sprintf("users = %v", users)
	return c.String(http.StatusOK, str)
}

func UpdateUser(c echo.Context) error {
	ctx := context.Background()
	db, err := connect(ctx)
	if err != nil {
		log.Println("connect failed.")
		log.Fatal(err.Error())
	}
	defer db.Close()

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	queries := sqls.New(db)
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	params := sqls.UpdateUserParams{
		ID:    id,
		Name:  u.Name,
		Email: u.Email,
	}

	queries.UpdateUser(ctx, params)
	if err != nil {
		log.Println("update user error.")
		return err
	}

	str := fmt.Sprintf("updated user id = %d", id)
	return c.String(http.StatusOK, str)
}

func DeleteUser(c echo.Context) error {
	ctx := context.Background()
	db, err := connect(ctx)
	if err != nil {
		log.Println("connect failed.")
		log.Fatal(err.Error())
	}
	defer db.Close()

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	queries := sqls.New(db)

	queries.DeleteUser(ctx, id)
	if err != nil {
		log.Println("delete user error.")
		return c.String(http.StatusNotFound, fmt.Sprintf("not found id = %d", id))
	}

	str := fmt.Sprintf("delete user id = %v", id)
	return c.String(http.StatusOK, str)
}

func connect(ctx context.Context) (*sql.DB, error) {
	db, err := sql.Open("mysql", "user:password@tcp(db:3306)/sample_db?charset=utf8mb4")
	if err != nil {
		return db, err
	}
	err = db.PingContext(ctx)
	if err != nil {
		return db, err
	}

	return db, err
}
