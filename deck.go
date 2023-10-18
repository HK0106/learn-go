package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	//cardNumber := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	cardNumber := []string{"2", "3", "4", "5", "6"}

	for i := 0; i < len(cardSuits); i++ {
		for j := 0; j < len(cardNumber); j++ {
			cards = append(cards, cardSuits[i]+" of "+cardNumber[j])
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
