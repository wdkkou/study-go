package main

import "fmt"

func main() {
	a1 := [3]int{1, 2, 3}
	a2 := [3]int{4, 5, 6}

	a1 = a2
	fmt.Printf("%v\n", a1)

	// それぞれの配列はメモリ上では別の領域のため、下の代入はa2には影響しない
	a1[0] = 0
	a1[2] = 0
	fmt.Printf("%v\n", a1) // => [0, 5, 0]
	fmt.Printf("%v\n", a2) // => [4, 5, 6]
}
