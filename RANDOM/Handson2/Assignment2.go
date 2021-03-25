package main

import "fmt"

const (
	c1     = 34
	c2 int = 89
)
const (
	a = 2017 + iota
	b = 2017 + iota
	c = 2017 + iota
)

var v int = 42
var str string = `raw: ,"KRISH", :string`

func main() {
	num := 45
	fmt.Printf("%d\t%b\t%x\n", num, num, num)
	fmt.Println(c1, c2)
	fmt.Printf("%d\t%b\t%#x\n", v, v, v)
	v2 := v << 1
	fmt.Printf("%d\t%b\t%#x\n", v2, v2, v2)
	fmt.Println(str)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
