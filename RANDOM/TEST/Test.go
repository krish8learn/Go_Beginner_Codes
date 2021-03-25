//this is 2nd sample program

package main

import (
	"fmt"
	"runtime"
)

const (
	a = iota
	b = iota
)

func main() {
	s := "LIVERPOOL"
	st := "krishnendu karmakar"
	fmt.Printf("Krish\n")
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println([]byte(s))
	fmt.Printf("%T\n", s)
	fmt.Println([]byte(st))
	fmt.Printf("%T\n", st)
	//converting the string utf-8 value into int
	abc := []byte(st)
	fmt.Printf("%T\n", abc) //unsigned integer with 8 bits size
	fmt.Println([]byte("A-Z"))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
		//fmt.Println()
		fmt.Printf("%U", s[i])
	}

	fmt.Println("iota example")
	fmt.Println(a + 10)
	fmt.Println(b)
	fmt.Printf("%T , %T \n", a, b)
}
