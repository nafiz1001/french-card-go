package card

import (
	"fmt"

	"github.com/nafiz1001/french-card-go/rank"
	"github.com/nafiz1001/french-card-go/suit"
)

type Card struct {
	Suit suit.Suit
	Rank rank.Rank
}

// String converts a Card to a string
func (c Card) String() string {
	return fmt.Sprintf("Card{%s,%s}", c.Suit, c.Rank)
}

// Rune converts a Card to its
// rune/unicode representation
func (c Card) Rune() rune {
	if _, err := CreateCard(c.Suit, c.Rank); err == nil {
		skipKnight := 0
		if c.Rank >= rank.Queen {
			skipKnight = 1
		}

		return rune(0x1F0A0 + int(c.Rank) + int(suit.Spade-c.Suit)*0x10 + skipKnight)
	} else {
		return rune(0)
	}
}

// CreateCard creates a Card if the suit s and rank r are valid
func CreateCard(s suit.Suit, r rank.Rank) (*Card, error) {
	if s < suit.Club || s > suit.Spade {
		return nil, fmt.Errorf("invalid suit number: %d", s)
	}

	if r < rank.Ace || r > rank.King {
		return nil, fmt.Errorf("invalid rank: %d", r)
	}

	return &Card{s, r}, nil
}

// CreateDeck creates a deck of 52 cards
func CreateDeck() []Card {
	cards := []Card{}

	for s := suit.Club; s <= suit.Spade; s++ {
		for r := rank.Ace; r <= rank.King; r++ {
			cards = append(cards, Card{s, r})
		}
	}

	return cards
}
