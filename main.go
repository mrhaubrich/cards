package main

func main() {
	cards := newDeckFromFile("my_cards")

	//hand, remainingDeck := deal(cards, 5)

	//fmt.Println("Mão")
	//hand.print()
	//fmt.Println("Deck")
	//remainingDeck.print()
	cards.shuffle()
	cards.print()

	//cards.saveToFile("my_cards")
}
