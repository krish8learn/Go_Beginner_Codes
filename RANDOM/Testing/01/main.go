package main

import (
	"fmt"
)

func main() {
	fmt.Println("2+3=", Mysum(2, 3))
	fmt.Println("3+4=", Mysum(3, 4))
	fmt.Println("1+9=", Mysum(1, 9))
}

func Mysum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
