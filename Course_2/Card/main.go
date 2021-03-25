package main

import "fmt"

func main() {
	var inhand deck
	var remain deck
	bunch := newDeck()
	bunch.shuffle()
	//bunch.show()
	for i := 0; i < 4; i++ {
		inhand, remain = deal(bunch, 5)
		fmt.Println("in hand of", i+1)
		inhand.show()
		inhand.savetofile("our_cards")
		//fmt.Println("Remains")
		//remain.show()
		bunch = remain
	}

}
