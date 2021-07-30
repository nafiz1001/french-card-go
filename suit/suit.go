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

type Suits []Suit

func (s Suits) Len() int           { return len(s) }
func (s Suits) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Suits) Less(i, j int) bool { return s[i] < s[j] }
