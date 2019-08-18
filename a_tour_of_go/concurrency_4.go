package main

import "fmt"

// CallFibonacci ...
func CallFibonacci() {
	ch := make(chan int, 10)
	go performFibonacci(ch, cap(ch))
	for i := range ch {
		fmt.Println(i)
	}
}

func performFibonacci(ch chan int, cbSize int) {
	x, y := 0, 1
	for _i := 0; _i < cbSize; _i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
