package main

import (
	"fmt"
	"time"
)

func main() {
	var sms = "Hello"
	go func() {
		fmt.Println(sms)
	}()
	sms = "NEW"
	go hello()
	time.Sleep(100 * time.Millisecond)
}

func hello() {
	fmt.Println("Inside hello")
}

//after launching the go routine main routine comes to next line, then the new go routine get the access of sms variable but by that time value of sms changed to hello to NEW
//to solve this use argument in go routines(arg)
//applying sleep is not good idea so we use waitgroup
//race check command ----> go run -race filename
