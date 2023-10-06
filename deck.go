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

package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

// (d deck) is the RECEIVEr
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
