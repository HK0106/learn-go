package main

import "fmt"

func main() {
	var card string = "Ace of Spades"
	card1 := "Ace of Heart"
	card2 := newCard()
	fmt.Println(card)
	fmt.Println(card1)
	fmt.Println(card2)

	//Array
	cards1 := []string{"A", "B", "C"}
	cards2 := append(cards1, "D")
	fmt.Println(cards1)
	fmt.Println(cards2)

	for i, card := range cards2 {
		fmt.Println(i, card)
	}

	//Slice
	fmt.Println(cards1)

}

func newCard() string {
	return "Five of Diamonds"
}
