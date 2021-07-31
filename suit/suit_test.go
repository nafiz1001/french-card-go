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
