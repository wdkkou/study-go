package main

import "fmt"

func createMap1() {

	m := make(map[int]string)
	m[1] = "US"
	m[81] = "Japan"
	m[86] = "China"

	fmt.Println(m)
}

func createMap2() {

	m := map[int]string{
		1: "Taro",
		2: "Hanako",
		3: "Jiro", // カンマが必要
	}

	fmt.Println(m)
}

func createMap3() {

	m := map[int][]int{
		1: {1},
		2: {1, 2},
		3: {1, 2, 3},
	}

	fmt.Println(m)
}

func referenceMap() {
	m := map[int]string{1: "A", 2: "B", 3: "C"}

	s1 := m[1]
	s2 := m[9]
	fmt.Println(s1)
	fmt.Println(s2) // => ""
}

func referenceMap2() {
	m := map[int]string{1: "A", 2: "B", 3: "C"}

	if _, ok := m[1]; ok {
		fmt.Println("m[1] is exist")
	}

	if _, ok := m[9]; !ok {
		fmt.Println("m[9] is not exist")
	}
}

func referenceMap3() {
	m := map[int][]int{
		1: {1},
		2: {1, 2},
		3: {1, 2, 3},
	}

	s := m[1]
	if s != nil {
		fmt.Println("m[1] is not nil")
	}
}

func forMap() {
	m := map[int]string{
		1: "Apple",
		2: "Banana",
		3: "Cherry",
	}

	for k, v := range m {
		fmt.Printf("%d => %s\n", k, v)
	}

}

func main() {
	// マップの生成
	createMap1()
	fmt.Println()
	createMap2()
	fmt.Println()
	createMap3()
	fmt.Println()

	referenceMap()
	fmt.Println()
	referenceMap2()
	fmt.Println()
	referenceMap3()
	fmt.Println()

	forMap()
	fmt.Println()
}
