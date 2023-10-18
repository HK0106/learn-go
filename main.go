package main

import "fmt"

func main() {
	cards := newDeck()

	cards.print()

	hand, remainCards := deal(cards, 3)

	fmt.Println(hand)
	fmt.Println(remainCards)
}
