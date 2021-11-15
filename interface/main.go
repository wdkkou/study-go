package main

import "fmt"

type MyError struct {
	Message string
	ErrCode int
}

// errorインターフェースのメソッドを実装
func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "error has occurred.", ErrCode: 999}
}

func callRaiseError() {
	err := RaiseError()
	fmt.Println(err.Error())
	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrCode)
	}

}

type Stringify interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("%s(%d)", p.Name, p.Age)
}
