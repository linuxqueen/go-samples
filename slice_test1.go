package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := a[:]
	b[1] = 9
	b = append(b, 8)
	for i, v := range b {
		fmt.Println(i, ":", v)
	}

}
