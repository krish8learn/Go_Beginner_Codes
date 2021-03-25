package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"liverpool", "Chelsea", "West Ham", "Aston Villa"},
	}
	bs, erv := toJSON(p1)
	if erv != nil {
		fmt.Println(erv)
		return
	}
	fmt.Println(string(bs))
}
func toJSON(p person) ([]byte, error) {
	bs, err := json.Marshal(p)
	if err != nil {
		return []byte{}, fmt.Errorf("Error: %v", err)
		//return []byte{},errors.New(fmt.Sprintf("Error: %v",err))
	}
	return bs, nil
}

//1st one returns error message as error type
//2nd one convert the error into string using fmt.Sprint() and errors.New() converts the string into error type
