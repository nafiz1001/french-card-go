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

func (c *Card) String() string {
	return fmt.Sprintf("Card{%s,%s}", c.Suit, c.Rank)
}

func (c *Card) Rune() rune {
	if _, err := CreateCard(c.Suit, c.Rank); err == nil {
		return rune(0x1F0A1 + int(c.Rank) + int(suit.Spade-c.Suit)*0x10)
	} else {
		return rune(0)
	}
}

func CreateCard(s suit.Suit, r rank.Rank) (*Card, error) {
	if s < suit.Club || s > suit.Spade {
		return nil, fmt.Errorf("invaid suit number: %d", s)
	}

	if r < rank.Ace || r > rank.King {
		return nil, fmt.Errorf("invalid rank: %d", r)
	}

	return &Card{s, r}, nil
}

func CreateDeck() []Card {
	cards := []Card{}

	for s := suit.Club; s <= suit.Spade; s++ {
		for r := rank.Ace; r <= rank.King; r++ {
			cards = append(cards, Card{s, r})
		}
	}

	return cards
}
