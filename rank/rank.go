package rank

import "strconv"

type Rank int

const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

func (r Rank) String() string {
	rankToStr := map[Rank]string{
		Ace:   "Ace",
		Jack:  "Jack",
		Queen: "Queen",
		King:  "King",
	}

	if Ace < r && r < Jack {
		return strconv.Itoa(int(r))
	} else {
		return rankToStr[r]
	}
}
