package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type employee struct {
	EmpID     string `json:"empid"`
	EmpName   string `json:"empname"`
	EmpMail   string `json:"empmail"`
	EmpMobile string `json:"empmobile"`
}
type coasterHandlers struct {
	store map[string]employee
}

func (h *coasterHandlers) get(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside get")
	//we are reading the data from the request, this will be printed on terminal
	data, er := ioutil.ReadAll(r.Body)
	if er != nil {
		http.Error(w, "oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, " Here is the Request received data %s\n", data)

	var details employee

	for _, coaster := range h.store {
		res := bytes.Compare(data, []byte(coaster.EmpID))
		if res == 0 {
			details = coaster
			break
		}
	}

	jsonBytes, err := json.Marshal(details)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}
func newCoasterHandlers() *coasterHandlers {
	return &coasterHandlers{
		store: map[string]employee{
			"first": employee{
				EmpID:     "1",
				EmpName:   "torres",
				EmpMail:   "fer@gmail.com",
				EmpMobile: "234999956677",
			},
			"second": employee{
				EmpID:     "2",
				EmpName:   "Casillas",
				EmpMail:   "iker@gmail.com",
				EmpMobile: "9598623777",
			},
			"thrird": employee{
				EmpID:     "3",
				EmpName:   "krish",
				EmpMail:   "krish@gmail.com",
				EmpMobile: "3234198577",
			},
			"Fourth": employee{
				EmpID:     "4",
				EmpName:   "Dortmund",
				EmpMail:   "bvb@gmail.com",
				EmpMobile: "23456677",
			},
		},
	}

}
func main() {
	fmt.Println("server active")
	fmt.Println("Enter ID:")
	coasterHandlers := newCoasterHandlers()
	http.HandleFunc("/coasters/", coasterHandlers.get)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
