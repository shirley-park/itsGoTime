package main

import "fmt"

func main() {
	/* VARIABLE DECLARATIONS
	------------------------
	var card string = "Ace of Spades" -> long form way of declaring variables
	card := "Ace of Spades" -> Go can figure out that card is a string
	card = "Five of Hearts" -> Reassigning card to "Five of Hearts"
	*/

	/* FUNCTIONS AND RETURN TYPES
	instead of declaring a var, calling a function that will return the card
	------------------------
	*/
	card := newCard()

	fmt.Println(card)

}

// defining a separate function that is called from main().
// whenever newCard() is called, it will return the card value of type string
func newCard() string {
	return "test"
}
