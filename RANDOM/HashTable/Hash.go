package main

import (
	"fmt"
)

func hash(word string, buckets int) int {
	letter := int(word[0])
	fmt.Println(letter)
	letter1 := rune(word[0])
	fmt.Println(letter1)
	//ret2 := letter1 % buckets
	ret := letter % buckets
	return ret
}

func main() {
	letter := 'A' //letter:= rune("A"[0])
	fmt.Println(letter)
	fmt.Printf("%T\n", letter)
	for i := 65; i < 123; i++ {
		fmt.Println(i, "---", string(i), "---", i%12)
	}
	n := hash("Go", 12)
	fmt.Println(n)
}
