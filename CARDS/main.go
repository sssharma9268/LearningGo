package main

func main() {
	//var card string = "Ace of Spades"
	//card := "Ace of Spades"  Another way to delcare a variable with its type
	//card = "Five of Diamonds" Colon is used only first time while declaring
	// card := newCard()
	//cards := []string{"Ace of Diamonds", newCard()}
	//cards := deck{"Ace of Diamonds", newCard()}
	//cards = append(cards, "Six of Spades") //appending new element to the slice
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	//cards.print()
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	//fmt.Println(cards)

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

}
