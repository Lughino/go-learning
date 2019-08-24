package main

// create a new type of desk chich is a slice of string
type deck []string

func (d deck) print() {
	for i, card := range d {
		println(i, card)
	}
}
