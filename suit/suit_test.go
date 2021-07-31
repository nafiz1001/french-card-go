package suit

import (
	"testing"
)

func TestSuitString(t *testing.T) {
	var s Suit

	s = Club
	if s.String() != "Club" {
		t.Errorf("expected Club.String() to return 'Club' but it returned '%s'", s)
	}

	s = Diamond
	if s.String() != "Diamond" {
		t.Errorf("expected Diamond.String() to return 'Diamond' but it returned '%s'", s)
	}

	s = Heart
	if s.String() != "Heart" {
		t.Errorf("expected Heart.String() to return 'Heart' but it returned '%s'", s)
	}

	s = Spade
	if s.String() != "Spade" {
		t.Errorf("expected Spade.String() to return 'Spade' but it returned '%s'", s)
	}

	s = Spade + 1
	if s.String() != "" {
		t.Errorf("expected (Spade + 1).String() to return '' but it returned '%s'", s)
	}

	s = Club - 1
	if s.String() != "" {
		t.Errorf("expected (Club - 1).String() to return '' but it returned '%s'", s)
	}
}
