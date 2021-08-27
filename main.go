package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards.print()
	//cards := newDeck()
	//cards.saveToFile("my_cards.txt")
	// hand, remainingDeck := deal(cards, 5)

	// cards.print()
	// hand.print()
	// remainingDeck.print()
}
