package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
	//hand, remainingDeck := deal(cards, 5)

	//fmt.Println("Mão")
	//hand.print()
	//fmt.Println("Deck")
	//remainingDeck.print()

}

func newCard() string {
	return "Three of Diamonds"
}
