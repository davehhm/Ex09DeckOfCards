package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{R: Ace, S: Spade})
	fmt.Println(Card{R: Three, S: Heart})
	fmt.Println(Card{R: Ten, S: Club})
	fmt.Println(Card{R: Queen, S: Diamond})
	fmt.Println(Card{R: Joker})

	// Output:
	// Ace_Spade
	// Three_Heart
	// Ten_Club
	// Queen_Diamond
	// Joker
}

func TestNewDeck(t *testing.T) {
	d := New(3)

	if len(d) != 13*4*3 {
		t.Errorf("Expected %d cards in 3 decks, returned %d", 13*4*3, len(d))
	}
}
