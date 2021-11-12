package main

import "fmt"

func main() {
	fmt.Println(plus(1, 2))
	fmt.Println(div(19, 7))

	f := returnFunc()
	f()
	returnFunc()()

	callFunction(func() {
		fmt.Println("I'm a function")
	})
}

func plus(x, y int) int {
	return x + y
}

func div(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}

func returnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func callFunction(f func()) {
	f()
}
