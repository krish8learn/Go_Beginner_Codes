package main

import "fmt"

var x int = 45
var y string = "JAMES BOND"
var z bool = true

type footballer string

var abc footballer

type imagine int

var jk imagine
var ran int

func main() {
	bit := 1          //0 or 1 or 2^0
	bytes := bit * 8  //2^3
	kb := bytes * 128 //2^10
	mb := kb * 1024
	gb := mb * 1024
	fmt.Printf("%d\t\t%b \n", bit, bit)
	fmt.Printf("%d\t\t%b \n", bytes, bytes)
	fmt.Printf("%d\t\t%b \n", kb, kb)
	fmt.Printf("%d\t\t%b \n", mb, mb)
	fmt.Printf("%d\t=%b \n", gb, gb)
	assignment1()
	assignment2()
	assignment4()
	assignment5()
	fmt.Println("INSIDE THE ASSIGNMENT3")
	//converting the values into string then printing them
	s := fmt.Sprintf("%v,%v,%v \n", x, y, z)
	fmt.Println(s)
}

func assignment1() {
	fmt.Println("INSIDE THE ASSIGNMENT1")
	x := 43
	y := "BOND"
	z := true
	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func assignment2() {
	fmt.Println("INSIDE THE ASSIGNMENT2")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}

func assignment4() {
	fmt.Println("INSIDE ASSIGNMENT 4")
	fmt.Println(abc)
	fmt.Printf("TYPE: ")
	fmt.Printf("%T\n", abc)
	abc = "CR7"
	fmt.Println(abc)

}

func assignment5() {
	ran = int(jk)
	fmt.Printf("%T", ran)
	fmt.Println(ran)
}
