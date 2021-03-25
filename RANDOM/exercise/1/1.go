package main

import (
	"fmt"
)

func main() {
	tot := func(n ...int) int {
		var largest int
		for _, v := range n {
			if v > largest {
				largest = v
			}
		}
		return largest
	}(4, 71, 4, 2, 45, 23, 12, 34, 96)

	fmt.Println(tot)
}
