package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	//card := "Ace of Spades" //compilers reads out that it is string
	// := should only be used while intialization
	//card = "Five of Diamonds"
	//card := newCard()
	cards := newDeck() //deck{"Ace of diamonds", newCard()}
	//cards = append(cards, "Six of Spades")
	//fmt.Println(cards.toString())
	cards.saveToFile("my_cards")
	cards.shuffle()
	fmt.Println("Shuffled Cards")
	cards.print()
	hand, remainingDeck := deal(cards, 5)
	fmt.Println("Hand of Cards")
	hand.print()
	fmt.Println("Remaining Cards")
	remainingDeck.print()
}
