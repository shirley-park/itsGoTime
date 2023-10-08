package main

/* VARIABLE DECLARATIONS

var card string = "Ace of Spades" -> long form way of declaring variables
card := "Ace of Spades" -> Go can figure out that card is a string
card = "Five of Hearts" -> Reassigning card to "Five of Hearts"
fmt.Println(card)
*/

// ------------------------

/* FUNCTIONS AND RETURN TYPES
instead of declaring a var, calling a function that will return the card

card := newCard()
fmt.Println(card)
*/

// ------------------------

/* SLICES AND FOR LOOPS
Go has two basic data structures for handling lists of records
> Array:
	- fixed length list of things
> Slice:
	- An array that can grow or shrink, can add and subtract new cards out of slice
	- Every element in a slice must be of same type (ie. all string or int)

cards := deck{"Ace of Diamonds", newCard()}
cards = append(cards, "Six of Spades")
// note: append() does not modify the existing slice
// instead returns a new slice which we assign back to var cards

how do we iterate over the cards slice?

	for i, card := range cards {
		fmt.Println(i, card)
	}
*/

func main() {
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")
	// cards.print()

	cards := newDeck()

	// deal returns two values
	// the first will be assigned to hand
	// second return value assigned to remainingCards
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}

// defining a separate function that is called from main().
// whenever newCard() is called, it will return the card value of type string
//- - - - - - - - - - - - -
// func newCard() string {
// 	return "Five of Diamonds"
// }
