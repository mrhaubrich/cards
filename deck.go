package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
		"Jack", "Queen", "King"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
func newDeckFromFile(filename string) (deck, error) {
	bs, err := ioutil.ReadFile(filename)

}