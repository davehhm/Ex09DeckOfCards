// Package main
package main

import (
	"fmt"

	"github.com/davehhm/Ex09DeckOfCards/deck"
)

func main() {
	cards := deck.New(2,
		deck.WithSort(deck.DefaultLess),
		deck.WithShuffle(),
		deck.WithJokers(5),
		//deck.WithFilter([]deck.Card{{R: deck.Ace}, {S: deck.Diamond}}),
	)
	/* sort.Slice(cards, func(i, j int) bool {
		return (cards[i].R < cards[j].R) ||
			((cards[i].R == cards[j].R) && (cards[i].S < cards[j].S))
	}) */
	fmt.Println(cards)
}
