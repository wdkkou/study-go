package controller

import (
	"context"
	"database/sql"
	"log"
	sqls "sample-api/app/db"
)

func CreateUser() (result sql.Result, err error) {
	ctx := context.Background()
	db, err := connect(ctx)
	if err != nil {
		log.Println("connect failed.")
		log.Fatal(err.Error())
	}
	defer db.Close()
	queries := sqls.New(db)

	return queries.CreateUser(ctx, sqls.CreateUserParams{
		Name:  "koki wada",
		Email: sql.NullString{String: "xxxx@yyyy.com", Valid: true},
	})
}

func GetUser(id int64) (user sqls.User, err error) {
	ctx := context.Background()
	db, err := connect(ctx)
	if err != nil {
		log.Println("connect failed.")
		log.Fatal(err.Error())
	}
	defer db.Close()
	queries := sqls.New(db)

	return queries.GetUser(ctx, id)
}

func GetUsers() (users []sqls.User, err error) {
	ctx := context.Background()
	db, err := connect(ctx)
	if err != nil {
		log.Println("connect failed.")
		log.Fatal(err.Error())
	}
	defer db.Close()
	queries := sqls.New(db)

	return queries.ListUsers(ctx)
}

func DeleteUser(id int64) error {
	ctx := context.Background()
	db, err := connect(ctx)
	if err != nil {
		log.Println("connect failed.")
		log.Fatal(err.Error())
	}
	defer db.Close()
	queries := sqls.New(db)

	return queries.DeleteUser(ctx, id)
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
