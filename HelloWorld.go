package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pradipta is Billa")
	//mapTryNoElse("afd")

	//http.Handle("/metrics", promhttp.Handler())
	//http.ListenAndServe(":2112", nil)
	cards := deck{}

	cardsShape := []string {"Spade", "Club"}
	cardsNumber := []string {"Ace", "King"}

	for i, shape := range cardsShape{
		for j, number := range cardsNumber{
			if i == j{

			}
			var now string = number+" "+shape
			cards = append(cards, now)
		}
	}
	cards.print()

}
