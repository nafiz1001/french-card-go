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
