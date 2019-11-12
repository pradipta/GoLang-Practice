package main

import (
	"fmt"
	"strings"
)

type deck []string

func newDeck() deck{

	cards := deck{}
	cardsShape := []string {"Spade", "Club"}
	cardsNumber := []string {"Ace", "King"}
	//cardsRange := cardsNumber[0:1]

	for _, shape := range cardsShape{
		for _, number := range cardsNumber{
			var now string = number+" "+shape
			cards = append(cards, now)
		}
	}
	return cards
}

func (d deck) print(){
	for i, card := range d {
		if i<0{
			i++
		}
		fmt.Println(card)
	}
}

func deal (d deck, size int) (deck, deck){
	 return d[:size], d[size:]
}

func (d deck) toString() string{
	return strings.Join([]string (d), ",")
}

func (d deck) saveToFile() {

}