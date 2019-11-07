package main

import "fmt"

type deck []string

func (d deck) print(){
	for i, card := range d {
		if i<0{
			i++
		}
		fmt.Println(card)
	}
}