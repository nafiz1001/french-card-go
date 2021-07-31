package main

import (
	"fmt"

	"github.com/nafiz1001/french-card-go/card"
	"github.com/nafiz1001/french-card-go/suit"
)

func main() {
	for s := suit.Club; s <= suit.Spade; s++ {
		fmt.Printf("%s\t%c\n", s, s.Rune())
	}

	fmt.Println("")
	cards := card.CreateDeck()
	for _, c := range cards {
		fmt.Printf("%-16s\t%c\n", c, c.Rune())
	}
}
