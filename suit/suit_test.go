package suit

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func TestSuitString(t *testing.T) {
	suitsToStr := map[Suit]string{
		Club:    "Club",
		Diamond: "Diamond",
		Heart:   "Heart",
		Spade:   "Spade",
	}

	suitsToName := suitsToStr

	for s := range suitsToStr {
		if s.String() != suitsToStr[s] {
			t.Errorf("expected %s.String() to return '%s' but it returned '%s'", suitsToName[s], suitsToStr[s], s)
		}
	}

	var s Suit

	s = Spade + 1
	if s.String() != "" {
		t.Errorf("expected (Spade + 1).String() to return '' but it returned '%s'", s)
	}

	s = Club - 1
	if s.String() != "" {
		t.Errorf("expected (Club - 1).String() to return '' but it returned '%s'", s)
	}
}

func TestSuitRune(t *testing.T) {
	suitsToName := []string{
		"Club",
		"Diamond",
		"Heart",
		"Spade",
	}

	file, err := os.Open("suits.txt")
	if err != nil {
		log.Fatal("could not open to read cards.txt")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for i, n := range suitsToName {
		s := Suit(i)
		if !scanner.Scan() {
			log.Fatal("ran out of cards to test against")
		}

		want := []rune(scanner.Text())[0]
		if s.Rune() != want {
			t.Errorf("expected %s.Rune() to return '%x' but it returned '%x'", n, want, s.Rune())
		}
	}

	var s Suit

	s = Spade + 1
	if s.Rune() != '\u0000' {
		t.Errorf("expected (Spade + 1).Rune() to return '\u0000' but it returned '%x'", s.Rune())
	}

	s = Club - 1
	if s.Rune() != '\u0000' {
		t.Errorf("expected (Club - 1).Rune() to return '\u0000' but it returned '%x'", s.Rune())
	}
}
