package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pradipta is Billa")
	//mapTryNoElse("afd")

	//http.Handle("/metrics", promhttp.Handler())
	//http.ListenAndServe(":2112", nil)

	cards := newDeck()

	cards.print()

}
