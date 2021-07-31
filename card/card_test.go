package card

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		{suit.Spade + 1, rank.Ace},
		{suit.Club - 1, rank.Ace},
		{suit.Club, rank.King + 1},
		{suit.Club, rank.Ace - 1},
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

	invalidArgs := []Card{
		{suit.Spade + 1, rank.Ace},
		{suit.Club - 1, rank.Ace},
		{suit.Club, rank.King + 1},
		{suit.Club, rank.Ace - 1},
	}

	var c Card

	c = invalidArgs[0]
	expectation := "Card{,Ace}"
	if c.String() != expectation {
		t.Errorf("expected c.String() to return %s but it returned %s", expectation, c.String())
	}

	c = invalidArgs[1]
	expectation = "Card{,Ace}"
	if c.String() != expectation {
		t.Errorf("expected c.String() to return %s but it returned %s", expectation, c.String())
	}

	c = invalidArgs[2]
	expectation = "Card{Club,}"
	if c.String() != expectation {
		t.Errorf("expected c.String() to return %s but it returned %s", expectation, c.String())
	}

	c = invalidArgs[3]
	expectation = "Card{Club,}"
	if c.String() != expectation {
		t.Errorf("expected c.String() to return %s but it returned %s", expectation, c.String())
	}
}

func TestCardRune(t *testing.T) {
	cards := CreateDeck()

	file, err := os.Open("cards.txt")
	if err != nil {
		log.Fatal("could not open to read cards.txt")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for _, c := range cards {
		if !scanner.Scan() {
			log.Fatal("ran out of cards to test against")
		}

		want := []rune(scanner.Text())[0]
		if c.Rune() != want {
			t.Errorf("expected c.Rune() to return %x but it returned %x", want, c.Rune())
		}
	}

	invalidArgs := []Card{
		{suit.Spade + 1, rank.Ace},
		{suit.Club - 1, rank.Ace},
		{suit.Club, rank.King + 1},
		{suit.Club, rank.Ace - 1},
	}

	for _, c := range invalidArgs {
		if c.Rune() != rune(0) {
			t.Errorf("expected c.Rune() to return %x but it returned %x", rune(0), c.Rune())
		}
	}
}
