package card

import (
	"fmt"
	"testing"

	"github.com/nafiz1001/french-card-go/rank"
	"github.com/nafiz1001/french-card-go/suit"
)

func TestCreateCard(t *testing.T) {
	for s := suit.Club; s <= suit.Spade; s++ {
		for r := rank.Ace; r <= rank.King; r++ {
			if c, err := CreateCard(s, r); c == nil || err != nil {
				t.Errorf("expected to successfully create a card with suit %s and rank %s", s, r)
			} else if c.Suit != s || c.Rank != r {
				t.Errorf("expected to create Card{%s,%s}, but created Card{%s,%s}", s, r, c.Suit, c.Rank)
			}
		}
	}

	invalidArgs := []Card{
		Card{suit.Spade + 1, rank.Ace},
		Card{suit.Club - 1, rank.Ace},
		Card{suit.Club, rank.King + 1},
		Card{suit.Club, rank.Ace - 1},
	}

	for i := range invalidArgs {
		s, r := invalidArgs[i].Suit, invalidArgs[i].Rank
		if c, err := CreateCard(s, r); c != nil || err == nil {
			t.Errorf("expected to fail create a card with suit #%d and rank #%d", s, r)
		}
	}
}

func TestCreateDeck(t *testing.T) {
	cards := CreateDeck()

	for s := suit.Club; s <= suit.Spade; s++ {
		for r := rank.Ace; r <= rank.King; r++ {
			index := int(s)*int(rank.King) + int(r-1)
			if cards[index].Suit != s || cards[index].Rank != r {
				t.Errorf("expected %s at index %d but found %s", Card{s, r}, index, cards[index])
			}
		}
	}
}

func TestCardString(t *testing.T) {
	cards := CreateDeck()

	for _, c := range cards {
		expectation := fmt.Sprintf("Card{%s,%s}", c.Suit, c.Rank)
		if c.String() != expectation {
			t.Errorf("expected c.String() to return %s but it returned %s", expectation, c.String())
		}
	}
}
