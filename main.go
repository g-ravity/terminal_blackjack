package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

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
	fmt.Println("Hello, Welcome to the game of Blackjack!")
	fmt.Println("I am the dealer.")

	cardsDeck := createDeck()
	cardsDeck.shuffle()

	numToDeal := 2

	dealerCards, cardsDeck := cardsDeck.deal(numToDeal)
	playerCards, cardsDeck := cardsDeck.deal(numToDeal)

	// dealer := Player{
	// 	faceUp:   Deck{dealerCards[0]},
	// 	faceDown: Deck{dealerCards[1]},
	// 	total:    dealerCards[0].num[0],
	// }

	fmt.Println("Here are my cards: ")
	for i := 0; i < numToDeal; i++ {
		if i == numToDeal-1 {
			PrintCard(dealerCards[i], false)
		} else {
			PrintCard(dealerCards[i], true)
		}
	}

	fmt.Println("Here are your cards: ")
	for i := 0; i < numToDeal; i++ {
		PrintCard(playerCards[i], true)
	}
}

func (v Value) getNum() []int {
	if v == Ace {
		return []int{1, 11}
	} else if v == King || v == Queen || v == Joker {
		return []int{10}
	} else {
		num, err := strconv.Atoi(string(v))

		if err != nil {
			fmt.Println("Error: ", err)
			return []int{0}
		}

		return []int{num}
	}
}

func createDeck() Deck {
	cards := Deck{}

	for _, suit := range []Suit{Hearts, Diamonds, Spades, Clubs} {
		for _, value := range []Value{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, King, Queen, Joker} {
			cards = append(cards, Card{suit: suit, value: value, num: value.getNum()})
		}
	}

	return cards
}

func (d Deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d Deck) deal(num int) (Deck, Deck) {
	return d[:num], d[num:]
}

func PrintCard(card Card, faceUp bool) {
	width := 9
	height := 7

	fmt.Println("┌─" + strings.Repeat("─", width-2) + "─┐")

	if faceUp {
		fmt.Printf("│ %-"+fmt.Sprintf("%d", width-4)+"s   │\n", card.value)

		fmt.Println("│ " + strings.Repeat(" ", width-2) + " │")

		fmt.Printf("│  %s%s%s │\n",
			strings.Repeat(" ", (width-4)/2),
			card.suit,
			strings.Repeat(" ", (width-3)/2))

		fmt.Println("│ " + strings.Repeat(" ", width-2) + " │")

		spacing := ""
		if len(card.value) <= 1 {
			spacing = " "
		}

		fmt.Printf("│%s %s%-1s │\n", strings.Repeat(" ", width-4), spacing, card.value)
	} else {
		for i := 1; i < height-1; i++ {
			if i == 3 {
				fmt.Println("│   ♦ ♣   │")
			} else {
				fmt.Println("│   ███   │")
			}
		}
	}

	fmt.Println("└─" + strings.Repeat("─", width-2) + "─┘")
}
