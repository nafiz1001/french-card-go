package rank

import (
	"testing"
)

func TestRankString(t *testing.T) {
	rankToName := map[Rank]string{
		Ace:   "Ace",
		Two:   "Two",
		Three: "Three",
		Four:  "Four",
		Five:  "Five",
		Six:   "Six",
		Seven: "Seven",
		Eight: "Eight",
		Nine:  "Nine",
		Ten:   "Ten",
		Jack:  "Jack",
		Queen: "Queen",
		King:  "King",
	}

	rankToStr := map[Rank]string{
		Ace:   "Ace",
		Two:   "2",
		Three: "3",
		Four:  "4",
		Five:  "5",
		Six:   "6",
		Seven: "7",
		Eight: "8",
		Nine:  "9",
		Ten:   "10",
		Jack:  "Jack",
		Queen: "Queen",
		King:  "King",
	}

	for rank := range rankToName {
		if rank.String() != rankToStr[rank] {
			t.Errorf("expected %s.String() to return '%s' but it returned '%s'", rankToName[rank], rankToStr, rank)
		}
	}

	var rank Rank

	rank = King + 1
	if rank.String() != "" {
		t.Errorf("expected (King + 1).String() to return '' but it returned '%s'", rank)
	}

	rank = Ace - 1
	if rank.String() != "" {
		t.Errorf("expected (Ace - 1).String() to return '' but it returned '%s'", rank)
	}
}
