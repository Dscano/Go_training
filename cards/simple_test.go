package main

import "fmt"

func main(){
	//var card string = "Ace of Spades" // I can write card:= "Ace of Spades"
	card := newCard()
	fmt.Println(card)

	cards := [] string {"Ace of Diamonds", newCard()} //slice definition, i.e dinamic array
	fmt.Println(cards)

	cards = append(cards,"Six of Spades")
	fmt.Println(cards)

	for i, card := range cards{ 
		fmt.Println(i, card)
	}

}

func newCard() string { // after() you can define the return tpe of the functions
	return "Five of Diamonds"
} 
