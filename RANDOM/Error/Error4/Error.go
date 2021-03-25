package main

import (
	"fmt"
	"log"
	"os"
)

//its provide us error message with date ,time
func main() {
	fmt.Println("Various Error message technique")
	_, err := os.Open("Krish.txt")
	if err != nil {
		fmt.Println("Error!!", err)
		log.Println("Error!!", err)
		log.Fatalln(err)
		log.Panicln(err)
		panic(err)
	}
}
