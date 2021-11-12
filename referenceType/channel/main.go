package main

import (
	"fmt"
	"time"
)

func reciever(ch <-chan int) {

	for {
		i := <-ch
		fmt.Println(i)
	}
}

func increment() {

	ch := make(chan int)

	go reciever(ch)

	i := 0

	for i < 10000 {
		ch <- i
		i++
	}

}

func printRune() {
	ch := make(chan rune, 3)

	ch <- 'A'
	ch <- 'B'
	ch <- 'C'
	ch <- 'D' // デッドロック
}

func receive(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		if ok == false {
			/* 受信できなくなったら終了*/
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + " is done.")

}

func functionGoroutine() {
	ch := make(chan int, 20)

	go receive("1st", ch)
	go receive("2nd", ch)
	go receive("3rd", ch)

	i := 0
	for i < 100 {
		ch <- i
		i++
	}
	close(ch)

	/* ゴールーチンの完了3秒待つ */
	time.Sleep(3 * time.Second)

}

func forChannel() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	for i := range ch {
		fmt.Println(i)
	}
}

func selectChannel() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)
	ch1 <- 1
	ch2 <- 2

	// 複数のcaseが成立するためランダムに選択される
	select {
	case <-ch1:
		fmt.Println("ch1から受信")
	case <-ch2:
		fmt.Println("ch2から受信")
	case ch3 <- 3:
		fmt.Println("ch3へ送信")
	default:
		fmt.Println("do not here.")
	}
}
func selectChannelAndGoroutine() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		for {
			i := <-ch1
			ch2 <- (i * 2)
		}
	}()

	go func() {
		for {
			i := <-ch2
			ch3 <- (i - 1)
		}
	}()

	n := 1

LOOP:
	for {
		select {
		case ch1 <- n:
			fmt.Println("ch1へ送信", n)
			n++
		case i := <-ch3:
			fmt.Println("ch3から受信", n)
			fmt.Println("recieved", i)
		default:
			fmt.Println("dafault", n)
			if n > 10 {
				break LOOP
			}
		}
	}
}

func main() {
	// increment()
	// printRune()
	// functionGoroutine()
	// forChannel()
	// selectChannel()
	selectChannelAndGoroutine()
}
