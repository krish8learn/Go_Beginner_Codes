// this is 1st go programme
package main

import "fmt"

var x = 43
var st = `krish say, "kroos his fav player"`

type footballer string

var Pname footballer

func main() {
	n, err := fmt.Println("hello")
	fmt.Println("bytes:", n)
	fmt.Println("error:", err)
	y := 98
	fmt.Println(x)
	fmt.Printf("TYPE: %T \n", x)
	fmt.Println(y)
	foo()
	fmt.Println(st)

	fmt.Printf("%b\n", x)
	fmt.Printf("%x\n", x)
	fmt.Printf("%#x\t %b\t %x\t %T\n", x, x, x, x)
	str := fmt.Sprintf("%#x\t %b\t %x\t %T\n", x, x, x, x)
	fmt.Println(str)

	conv()
}

func foo() {
	fmt.Println("Inside foo")
	z := "krish"
	fmt.Println(z)
	fmt.Printf("TYPE: %T \n", z)

	Pname = "TONI KROOS"
	fmt.Println(Pname)
	fmt.Printf("%T\n", Pname)
}

func conv() {
	//Pname which is footballer type has converted to string
	fmt.Printf("%T\n", Pname)
	abc := string(Pname)
	fmt.Println(abc)
	fmt.Printf("%T\n", abc)
}
