package suit

type Suit int

const (
	Club Suit = iota
	Diamond
	Heart
	Spade
)

func (s Suit) String() string {
	suitsToStr := map[Suit]string{
		Club:    "Club",
		Diamond: "Diamond",
		Heart:   "Heart",
		Spade:   "Spade",
	}

	return suitsToStr[s]
}

func (s Suit) Rune() rune {
	suitsToRune := map[Suit]rune{
		Club:    '\u2663',
		Diamond: '\u2666',
		Heart:   '\u2665',
		Spade:   '\u2660',
	}

	return suitsToRune[s]
}
