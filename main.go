package main

func main() {
	//var cards string = "Ace of spaces"
	cards := newDeck()

	cards.print()
	//fmt.Println(cards)

	hand, remainingHand := deal(cards, 5)
	hand.print()
	remainingHand.print()
}
