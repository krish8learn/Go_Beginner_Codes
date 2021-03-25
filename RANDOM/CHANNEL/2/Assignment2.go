package main

import (
	"fmt"
)

func main() {
	cs := make(chan<- int) //send channel

	go func() {
		cs <- 43
	}()

	//fmt.Println(<-cs) cannot perform receive operation
	fmt.Printf("-------\n")
	fmt.Printf("cs\t%T\n", cs)
}
