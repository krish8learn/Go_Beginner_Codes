package main

import (
	"errors"
	"fmt"
	"log"
)

type sqrtError struct {
	//lat  string
	//long string
	err error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("Math Error: %v", se.err)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, sqrtError{
			err: errors.New("-ve value"),
		}
	}
	return f * f, nil
}
func main() {
	data, err := sqrt(-6)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(data)
	}
}

//error.New() convert the string into error type
// we can also do fmt.Errorf("Error: %v",f) instead of error.New()
