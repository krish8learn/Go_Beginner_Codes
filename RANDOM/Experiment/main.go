package main

import (
	"RANDOM/Experiment/EXP"
	"fmt"
)

type canine struct {
	name string
	age  int
}

func main() {
	mambo := canine{
		name: "Mambo",
		age:  EXP.Years(4),
	}
	fmt.Println(mambo)
}
