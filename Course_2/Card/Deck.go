package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	var cards deck
	cardSuits := []string{"SPADES", "HEARTS", "CLUBS", "DIAMONDS"}
	cardValue := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "joker", "king", "queen"}
	for _, suits := range cardSuits {
		for _, values := range cardValue {
			cards = append(cards, values+" of "+suits)
		}
	}
	return cards
}

func (d deck) show() {
	for i, v := range d {
		fmt.Println(i, v)
	}
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	num := rand.New(source)
	for v := range d {
		newpos := num.Intn(len(d) - 1)
		d[v], d[newpos] = d[newpos], d[v]
	}
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) savetofile(filename string) error {
	if _, err := os.Stat(filename); err == nil {

	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}

	if data != nil {
		str := strings.Join([]string(d), ",")
		return ioutil.WriteFile(filename, []byte("\n"+str), 0666)
	}
	str := strings.Join([]string(d), ",")
	return ioutil.WriteFile(filename, []byte(str), 0666)

}
