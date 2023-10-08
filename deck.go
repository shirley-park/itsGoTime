/* THE GO APPROACH
- Go has the base Go types (string, integer, float, array, map, etc)
- We want to "extend" a base type and add some extra functionality to it
*/

// ------------------------

/*
> 'type deck[]string' tells Go we want to create an array of strings
- And add some functions specifically made to work with it

> Functions with 'deck' as a RECEIVER:
-> a function w a receiver is like a "method" - a function that belongs to an "instance"

>> So in the 'cards' folder we have:
-> main.go - code to create and manipulate a deck
-> deck.go - code that describes what a deck is & how it works
-> desk_test.go - code to automatically test the deck
*/

// ------------------------

/* SLICE RANGE SYNTAX
- Slices are zero-indexed:
eg. say we have a slice of fruits:
           0         1         2        3
fruits = "apple", "banana", "grape", "orange"

fruits[0] = "apple"
fruits[3] = "orange"

fruits[startIndexIncl : upToExcl]
fruits[0:2] = "apple", "banana"

we can optionally exclude numbers:
fruits[:2] = "apple", "banana"
meaning give me everything from beginning to index 2
OR
fruits[2:] = "grape", "orange"
*/

// ------------------------

package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

// newDeck() creates and returns a list of playing cards
func newDeck() deck {
	// cards of type deck
	// contains all 52 cards
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// (d deck) is the RECEIVER
// to any variable of type 'deck' get access to print() method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// we can call deal() with
// 1st arg of type deck and
// 2nd arg of type int
// (deck, deck) means this func will return two sep values of type deck
func deal(d deck, handSize int) (deck, deck) {

	// Multiple return values
	//take list of deck and split into 2 using handSize

	return d[:handSize], d[handSize:]

}
