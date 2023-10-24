// Package deck
package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type suit int16

const (
	_ suit = iota
	Spade
	Heart
	Club
	Diamond
)

func (s suit) String() string {
	switch s {
	case Spade:
		return "spade"
	case Heart:
		return "heart"
	case Club:
		return "club"
	case Diamond:
		return "diamond"
	}
	return "undefined"
}

type rank int16

const (
	_ rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Joker
)

func (r rank) String() string {
	switch r {
	case Ace:
		return "A"
	case Two:
		return "2"
	case Three:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case Ten:
		return "10"
	case Jack:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	case Joker:
		return "Joker"
	}
	return "undefined"
}

type Card struct {
	R rank
	S suit
}

// New a card deck, accepting options
func New(dcount int, options ...func([]Card) []Card) []Card {
	cards := []Card{}
	for i := 0; i < dcount; i++ {
		for s := Spade; s <= Diamond; s++ {
			for r := Ace; r <= King; r++ {
				cards = append(cards, Card{R: r, S: s})
			}
		}
	}

	for _, o := range options {
		cards = o(cards)
	}
	return cards
}

// Option of sorting a card deck. Taking a user provided "less" function for sorting
func WithSort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func DefaultLess(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return (cards[i].R < cards[j].R) ||
			((cards[i].R == cards[j].R) && (cards[i].S < cards[j].S))
	}
}

// Option of shuffling a card deck
func WithShuffle() func([]Card) []Card {
	return Shuffle()
}

// Shuffle a card deck
func Shuffle() func([]Card) []Card {
	return func(cards []Card) []Card {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		r.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
		return cards
	}
}

// Option of insert n numbers of Jokers into the deck
func WithJokers(n int) func([]Card) []Card {
	return func(c []Card) []Card {
		result := c
		// set seed
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 0; i < n; i++ {
			idx := r.Intn(len(result))
			card := Card{R: Joker}
			result = append(result[:idx], append([]Card{card}, result[idx:]...)...)
			//fmt.Println("resultlen:", len(result), " clen:", len(c))
			//fmt.Println(c)
			// WARNING: tested below will insert the card twice. append func affects the "c" variable
			//result = append(append(c[:idx], card), c[idx:]...)
			fmt.Println("resultlen:", len(result), " clen:", len(c))
			fmt.Println(c)
		}
		return result
	}
}

// Option to filter out specific cards
func WithFilter(fs []Card) func([]Card) []Card {
	return func(c []Card) []Card {
		var result []Card
		for _, v := range c {
			filtered := false
			for _, f := range fs {
				if f.R == 0 {
					if f.S == v.S {
						filtered = true
					}
				} else if f.S == 0 {
					if f.R == v.R {
						filtered = true
					}
				} else if f.S == v.S && f.R == v.R {
					filtered = true
				}
			}
			if !filtered {
				result = append(result, v)
			}
		}
		return result
	}
}
