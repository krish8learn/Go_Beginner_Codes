package main

import (
	"fmt"
)

func main() {
	var answer1, answer2 string

	fmt.Print("First Name:")
	_, err := fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Print("Last Name:")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)
	}

	fmt.Println()
	n, erv := fmt.Println(answer1, answer2)
	if erv != nil {
		fmt.Println("error")
		panic(erv)
	} else {
		fmt.Println(n)
	}

}
