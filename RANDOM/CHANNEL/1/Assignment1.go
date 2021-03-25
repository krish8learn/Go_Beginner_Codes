package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)

	c <- 34 //or we can use go with anonymous function

	fmt.Println(<-c)
}
