package main

import "fmt"

func main() {

	cards := newDeck()
	cards.print()

	
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

	fmt.Println(cards.toString())

	cards.saveToFile("my_cards")

	n_cards := newDeckFromFile("my_cards")
	n_cards.print()

	s_cards := newDeck()
	s_cards.shuffle()
	s_cards.print()

}

