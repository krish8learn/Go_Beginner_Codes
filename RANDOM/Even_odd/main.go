package main

import "fmt"

func main() {
	num := make([]int, 11)
	for i := 0; i < 11; i++ {
		num = append(num, i)
	}
	for index := range num {
		if num[index]%2 == 0 {
			fmt.Println(num[index], " is Even")
		} else if num[index]%2 != 0 {
			fmt.Println(num[index], " is Odd")
		}
	}
}
