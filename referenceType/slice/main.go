package main

import "fmt"

func main() {
	fmt.Println("--- スライスの拡張 ---")
	s := []int{1, 2, 3}
	s = append(s, 4)
	fmt.Println(s)

	s = append(s, 5, 6, 7)
	fmt.Println(s)

	sl := make([]int, 0, 0)
	fmt.Printf("len = %d, cap= %d\n", len(sl), cap(sl))

	sl = append(sl, 1)
	fmt.Printf("len = %d, cap= %d\n", len(sl), cap(sl))

	sl = append(sl, []int{2, 3, 4}...)
	fmt.Printf("len = %d, cap= %d\n", len(sl), cap(sl))
	fmt.Println()

	fmt.Println("--- forSlice関数 ---")
	forSlice()
	fmt.Println()

	fmt.Println("--- sum関数 ---")
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum())

	a := []int{1, 2, 3}
	fmt.Println(sum(a...))
	fmt.Println()

	fmt.Println("--- pow関数 ---")
	x := [3]int{1, 2, 3}
	// この関数は値渡しで、配列のコピーが行われているため、main関数内の変数xには影響されない
	pow(x)
	fmt.Println(x) // => [1, 2, 3]

	y := []int{1, 2, 3}
	// スライスを引数にすると、 参照渡しになり、main関数内の変数yの値が変わる
	powArgSlice(y)
	fmt.Println(y) // => [1, 4, 9]
	fmt.Println()

	fmt.Println("--- trapSlice関数 ---")
	trapSlice()
	fmt.Println()

}

func forSlice() {
	s := []string{"Apple", "Banana", "Cherry"}

	for i, v := range s {
		fmt.Printf("[%d] => %s\n", i, v)
	}

	// 要素を追加するとlen(s)が増え続け無限ループになる
	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("[%d] => %s\n", i, s[i])
	// 	s = append(s, "Melon")
	// }

	// 範囲節によるfor文だとスライスへの要素の追加はループ処理の回数に影響を与えない
	for i, v := range s {
		fmt.Printf("[%d] => %s\n", i, v)
		s = append(s, "Melon")
	}

	fmt.Println(s)
}

func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}

	return n
}

func pow(a [3]int) {
	for i, v := range a {
		a[i] = v * v
	}
	return
}

func powArgSlice(a []int) {
	for i, v := range a {
		a[i] = v * v
	}
	return
}

func trapSlice() {
	a := [3]int{1, 2, 3}
	s := a[:]
	fmt.Printf("len = %v\n", len(s))
	fmt.Printf("cap = %v\n", cap(s))
	s[1] = 0
	//　配列とスライスともに同じ配列のデータを共有
	fmt.Printf("a = %v\n", a)
	fmt.Printf("s = %v\n", s)

	fmt.Println()
	s = append(s, 4)
	fmt.Printf("len = %v\n", len(s))
	fmt.Printf("cap = %v\n", cap(s))
	s[1] = 100
	// スライスが自動拡張される場合、参照先が代わり、共有されない
	fmt.Printf("a = %v\n", a)
	fmt.Printf("s = %v\n", s)

}
