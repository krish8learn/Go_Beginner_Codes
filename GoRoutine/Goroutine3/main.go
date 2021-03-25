package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter int
var m = sync.RWMutex{}

func main() {
	fmt.Println("Inside main")
	for i := 0; i < 2; i++ {
		wg.Add(2)
		m.RLock()
		go hello()
		m.Lock()
		go hi()
	}
	wg.Wait()
	runtime.GOMAXPROCS(100)
	fmt.Println(runtime.NumCPU())
	fmt.Printf("THREADS: %v\n", runtime.GOMAXPROCS(4))
}
func hello() {
	fmt.Printf("Inside Hello line1\n")
	fmt.Println(counter, "Inside hello line2")
	m.RUnlock()
	wg.Done()
}
func hi() {
	fmt.Println("Inside hi line1")
	counter++
	fmt.Printf("Inside hi line2  %v\n", counter)
	m.Unlock()
	wg.Done()
}
