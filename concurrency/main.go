package main

import (
	"fmt"
	"runtime"
)

func sub() {
	for {
		fmt.Println("sub loop")
	}
}

func main() {
	// go sub()
	// for {
	// 	fmt.Println("main loop")
	// }

	go fmt.Println("Yeah!")

	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("NumGorutine: %d\n", runtime.NumGoroutine()) // go文によってgorutineが増える
	fmt.Printf("Version: %s\n", runtime.Version())
}
