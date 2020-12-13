package main

//you can initialize a var outside of a func, but you cannot assign anything to it
var testCard string

// test struct, does nothing
type hand struct {
	card1 string
	card2 string
	card3 string
	card4 string
	card5 string
}

func main() {

	// card := newCard()
	cards := newDeck()
	cards.shuffle()
	cards.print()

	// hand, remCards := deal(cards, 5)
	// hand.print()
	// remCards.print()

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// //declare variable and statically type it
	// var card2 string
	// //assign variable
	// card2 = "Two of Clubs"

	// //finally assigning test card
	// testCard = "Two of Spades"

	// playingCard := "King of Hearts" //variable declaration and assignment, only used when defining new var

	// //very longform method, don't really need to do this
	// var card string = "Queen of Spades" // only a string will ever be assigned to this variable (notifies compiler)

}

// new function that returns a string
// array -> fixed length
// slice -> varible length
// every element in array or slice must be of the same type
func newCard() string {
	return "Five of Diamonds"
}

//ExportCard exportable function (starts w/ capital letter)
func ExportCard() string {
	return "I can see you!"
}
