package main

import "fmt"

type Vector []int

const NCPU = 4

func DoSome(v Vector, c chan int) {
	result := 0
	fmt.Println("size=", len(v))
	for _, n := range v {
		result += n
	}
	c <- result
}

func main() {
	size := 1000000000
	v := make(Vector, size)
	all := 0
	for i := 1; i < size+1; i++ {
		v[i-1] = i
		all += i
	}
	fmt.Println("all=", all)

	c := make(chan int, NCPU)

	step := size/NCPU + 1
	fmt.Println("step=", step)
	i := 0
	for ; i < NCPU-1; i++ {
		go DoSome(v[i*step:(i+1)*step], c)
	}
	go DoSome(v[i*step:], c)

	all = 0
	for i := 0; i < NCPU; i++ {
		x := <-c
		fmt.Println(x)
		all += x
	}

	fmt.Println("all=", all)
}
