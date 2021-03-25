package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("Random.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	bs, erv := ioutil.ReadAll(f)
	if erv != nil {
		fmt.Println(erv)
	}
	fmt.Println(string(bs))
}
