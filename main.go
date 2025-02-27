package main

import "fmt"

type Suit string

const (
	Hearts   Suit = "♥"
	Diamonds Suit = "♦"
	Spades   Suit = "♠"
	Clubs    Suit = "♣"
)

type Value string

const (
	Ace   Value = "A"
	Two   Value = "2"
	Three Value = "3"
	Four  Value = "4"
	Five  Value = "5"
	Six   Value = "6"
	Seven Value = "7"
	Eight Value = "8"
	Nine  Value = "9"
	Ten   Value = "10"
	King  Value = "K"
	Queen Value = "Q"
	Joker Value = "J"
)

type Card struct {
	suit  Suit
	value Value
	num   []int
}

type Deck []Card

type Player struct {
	faceUp   Deck
	faceDown Deck
	total    int
}

func main() {
	fmt.Println("Hello, World!")
}
