package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func (p *person) speak() {
	fmt.Println(p.Name, p.Age)
	fmt.Println("Footballer")
}

type human interface {
	speak()
}

func saysomething(h human) {
	h.speak()
}

func main() {
	fmt.Println("Inside main")
	p1 := person{
		Name: "Toni",
		Age:  31,
	}
	saysomething(&p1)
	p1.speak()
}
