package deck

import (
	"fmt"
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
