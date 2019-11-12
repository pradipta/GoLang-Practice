package main

import "fmt"

func main() {
	//fmt.Println("Pradipta is Billa")
	//mapTryNoElse("afd")

	//http.Handle("/metrics", promhttp.Handler())
	//http.ListenAndServe(":2112", nil)

	cards := newDeck()
	//cards.print()
	deck1, deck2 := deal (cards, 2)

	fmt.Println("--------")
	deck1.print()
	fmt.Println("--------")
	deck2.print()

	var byteCard []byte = []byte(cards[0])
	var deckToString string = cards.toString()

	fmt.Println(byteCard)
	fmt.Println(deckToString)
}
