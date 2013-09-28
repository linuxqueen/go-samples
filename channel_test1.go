package main

import "fmt"

func Count(index int, ch chan int) {
	fmt.Println("counting:", index)
	ch <- index
}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(i, chs[i])
	}

	for _, ch := range chs {
		fmt.Println("waiting ", <-ch)
	}
}
