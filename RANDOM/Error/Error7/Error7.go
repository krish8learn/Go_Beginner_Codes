package main

//Create struct customerr which implements in-built  error interface
import (
	"fmt"
)

type customerr struct {
	info string
}

func (ce customerr) Error() string {
	return fmt.Sprintf("ERROR OCCURED %v ", ce.info)
}
func foo(e error) {
	fmt.Println("Foo run:--", e, "\n")
}
func main() {
	c1 := customerr{
		info: "need more data",
	}
	foo(c1)
}
