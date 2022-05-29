[![codecov](https://codecov.io/gh/nafiz1001/french-card-go/branch/main/graph/badge.svg?token=ZSO0NKOMRI)](https://codecov.io/gh/nafiz1001/french-card-go)

# french-card-go

A set of structs and functions to represent [french-suited playing cards](https://en.wikipedia.org/wiki/French-suited_playing_cards)

# Features

## Suit
```go
const (
	Club Suit = iota
	Diamond
	Heart
	Spade
)
```

## Rank
```go
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
```

## Card
```go
type Card struct {
	Suit suit.Suit
	Rank rank.Rank
}
```
## Unicode

`go run cmd/print_unicodes.go`

```
Club    ♣
Diamond ♦
Heart   ♥
Spade   ♠

Card{Club,Ace}          🃑
Card{Club,2}            🃒
Card{Club,3}            🃓
Card{Club,4}            🃔
Card{Club,5}            🃕
Card{Club,6}            🃖
Card{Club,7}            🃗
Card{Club,8}            🃘
Card{Club,9}            🃙
Card{Club,10}           🃚
Card{Club,Jack}         🃛
Card{Club,Queen}        🃜
Card{Club,King}         🃝
...
Card{Spade,King}        🂭
```
