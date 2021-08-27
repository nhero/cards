package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	//cards := newDeck()
	//cards.saveToFile("my_cards.txt")
	// hand, remainingDeck := deal(cards, 5)

	// cards.print()
	// hand.print()
	// remainingDeck.print()
}
