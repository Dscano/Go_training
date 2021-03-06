package main

import "fmt"

func main() {

	cards := newDeck()
	fmt.Println("************* It is printing cards generated by newDeck() *************")
	cards.print()

	

	hand, remainingCards := deal(cards, 5)

	fmt.Println("************* It is printing cards in the hand *************")
	hand.print()
	fmt.Println("************* It is printing cards remained in a deck *************")
	remainingCards.print()
	fmt.Println("************* It is printing deks toString *************")
	fmt.Println(cards.toString())

	cards.saveToFile("my_cards")

	n_cards := newDeckFromFile("my_cards")
	fmt.Println("************* It is printing cards from a saved deck file *************")
	n_cards.print()

	s_cards := newDeck()
	s_cards.shuffle()
	fmt.Println("************* It is printing shuffled cards in a deck *************")
	s_cards.print()

}

