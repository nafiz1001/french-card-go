package suit

import (
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
	suitsToRune := map[Suit]rune{
		Club:    '\u2663',
		Diamond: '\u2666',
		Heart:   '\u2665',
		Spade:   '\u2660',
	}

	suitsToName := map[Suit]string{
		Club:    "Club",
		Diamond: "Diamond",
		Heart:   "Heart",
		Spade:   "Spade",
	}

	for s := range suitsToRune {
		if s.Rune() != suitsToRune[s] {
			t.Errorf("expected %s.Rune() to return '%x' but it returned '%x'", suitsToName[s], suitsToRune[s], s.Rune())
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
