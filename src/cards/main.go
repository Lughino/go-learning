package main

import "fmt"

var deckCount int

func main() {
	// var card = "Ace of Spades"
	card := newCard()
	deckCount = 5
	fmt.Println("deck count", deckCount)
	fmt.Println("simple card", card)

	// cards := []string{"Five of Diamonds"}
	cards := deck{"Five of Diamonds"}
	cards = append(cards, "Six of Spades", newCard())

	cards.print()

	// when index is not used, use a special "blank" character "_"
	for _, card := range cards {
		println("Without index", card)
	}
}

func newCard() string {
	return "Ace of Spades"
}
