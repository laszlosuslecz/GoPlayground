package main

import (
	"fmt"
)

var sum int //= 0

func main() {
	fmt.Println("Hello world")
	/*
		for i := 0; i < 10; i++ {
			sum++
			fmt.Println(sum)
		}
	*/
	for sum < 10 {
		sum++
		fmt.Println(sum)
	}
}
