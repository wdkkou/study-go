package model

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

type CreateUser struct {
	Id    int
	Name  string
	Email string
}

func New() *User {
	return &User{}
}
