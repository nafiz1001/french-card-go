package rank

import (
	"sort"
	"strconv"
	"testing"
)

func TestRankToString(t *testing.T) {
	var rank Rank

	rank = Ace
	if rank.String() != "Ace" {
		t.Errorf("expected Ace.String() to return 'Ace' but it returned '%s'", rank)
	}

	intToWord := map[Rank]string{
		Two:   "Two",
		Three: "Three",
		Four:  "Four",
		Five:  "Five",
		Six:   "Six",
		Seven: "Seven",
		Eight: "Eight",
		Nine:  "Nine",
		Ten:   "Ten",
	}

	for rank = Two; rank <= Ten; rank++ {
		if rank.String() != strconv.Itoa(int(rank)) {
			t.Errorf("expected %s.toString() to return '%d' but it returned '%s'", intToWord[rank], int(rank), rank)
		}
	}

	rank = Jack
	if rank.String() != "Jack" {
		t.Errorf("expected Jack.String() to return 'Ace' but it returned '%s'", rank)
	}

	rank = Queen
	if rank.String() != "Queen" {
		t.Errorf("expected Queen.String() to return 'Queen' but it returned '%s'", rank)
	}

	rank = King
	if rank.String() != "King" {
		t.Errorf("expected King.String() to return 'King' but it returned '%s'", rank)
	}

	rank = King + 1
	if rank.String() != "" {
		t.Errorf("expected (King + 1).String() to return '' but it returned '%s'", rank)
	}

	rank = Ace - 1
	if rank.String() != "" {
		t.Errorf("expected (Ace - 1).String() to return '' but it returned '%s'", rank)
	}
}

func TestRankSort(t *testing.T) {
	ranks := []Rank{}
	for r := King; r >= Ace; r-- {
		ranks = append(ranks, r)
	}

	sort.Sort(Ranks(ranks))

	for r := Ace; r <= King; r++ {
		if ranks[r-1] != r {
			t.Errorf("expected %s at %d but found %s", Rank(r), r, ranks[r])
		}
	}
}
