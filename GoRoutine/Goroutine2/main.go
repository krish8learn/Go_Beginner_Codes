package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{} //this actualy synchronize the new go routine and main go routine
func main() {
	var sms = "hello"
	wg.Add(1) //means we are adding one extra go routine
	go func(sms string) {
		fmt.Println(sms)
		wg.Done() //means extra go routine ended
	}(sms)
	sms = "NEW"
	wg.Wait() //telling main to wait for other go routine to execute
}
