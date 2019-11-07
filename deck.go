package main

import "fmt"

type deck []string

func newDeck() deck{

	cards := deck{}
	cardsShape := []string {"Spade", "Club"}
	cardsNumber := []string {"Ace", "King"}

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