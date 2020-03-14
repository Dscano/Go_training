package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
)

// Create a new type of 'deck', e.i It is a slices of strings
type deck []string

// Function for create a new deck
func newDeck() deck {
	cards := deck{}
	cardSuits := [] string{"Spades","Hearts","Diamonds","Clubs"}
	cardValues := [] string{"Ace","Two","Three","Four","Five","Six","Seven", "Eigth", "Nine", "Ten","Jack","Queen","King"}
    
    for _, suit := range cardSuits{
    	for _, value :=range cardValues{
    		cards = append(cards, value+" of "+suit)
    	}
    }

    return cards
}

// Function for printing the deck elements, it uses a reciver, ie (d deck)
func (d deck) print() {
	for i, card := range d{ 
		fmt.Println(i, card)
	}
}

// Function to extracts a certain number of cards from a deck
func deal(d deck, handSize int) (deck,deck) {
	return d [:handSize], d[handSize:]
}

//Function to transform a deck into a string
func (d deck) toString() string {
	return strings.Join([]string(d),",") 
}

//Function to save a deck into a file
func (d deck) saveToFile(filename string) error{
	// save a file with read a write priviledge
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
} 

//Function to read a deck from a file
func newDeckFromFile(filename string) deck {

	bs, err := ioutil.ReadFile(filename)
	if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
	}

	ss := strings.Split(string(bs),",")
	return deck(ss)
}

//Function to shuffle the cards 
func (d deck)shuffle () {
	// gerate a different int 64 number
	source := rand.NewSource(time.Now().UnixNano()) 
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// swap the elements at both i and newPosition inside the slice
		d[i], d[newPosition] = d[newPosition], d[i]  
	}

}