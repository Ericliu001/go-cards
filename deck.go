package main

type deck []string

func newDeck() deck {
	suits := []string{"diamonds", "clubs", "hearts", "spades"}
	ranks := []string{"ace", "king", "queen", "jack", "10", "9", "8", "7", "6", "5", "4", "3", "2"}

	var cards deck = deck{}

	for _, suit := range suits {
		for _, rank := range ranks {
			cards = append(cards, suit+" of "+rank)
		}
	}
	return cards
}

func (d deck) deal(hand int) (deck, deck) {
	return d[:hand], d[hand:]
}

func deal(d deck, hand int) (deck, deck) {
	return d[: hand], d[hand:]
}
