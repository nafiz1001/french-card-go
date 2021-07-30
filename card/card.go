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

type Cards []Card

func (c Cards) Len() int      { return len(c) }
func (c Cards) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c Cards) Less(i, j int) bool {
	return c[i].Suit < c[j].Suit || (c[i].Suit == c[j].Suit && c[i].Rank < c[j].Rank)
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

func CreateDeck() Cards {
	cards := Cards{}

	for s := suit.Club; s <= suit.Spade; s++ {
		for r := rank.Ace; r <= rank.King; r++ {
			c, _ := CreateCard(s, r)
			cards = append(cards, *c)
		}
	}

	return cards
}
