package main

import (
	"fmt"
	"runtime"
	"sync"
	//"time"
)

func myfunc1(wg *sync.WaitGroup) {
	//time.Sleep(1 * time.Second)
	fmt.Println("Inside myfunc1")
	wg.Done()
}

func myfunc2(wg *sync.WaitGroup) {
	//time.Sleep(1 * time.Second)
	fmt.Println("Inside myfunc2")
	wg.Done()
}

func main() {
	fmt.Println("Inside main")
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)
	go myfunc1(&wg)
	fmt.Println(runtime.NumGoroutine())
	go myfunc2(&wg)
	fmt.Println(runtime.NumGoroutine())
	wg.Wait() //hold the main function so that two myfunc can finish
	fmt.Println(runtime.NumGoroutine())
	fmt.Println("Exiting main")
	fmt.Println(runtime.NumGoroutine())
}
