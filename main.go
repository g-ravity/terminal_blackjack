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

	dealer := Player{
		faceUp:   dealerCards[:len(dealerCards)-1],
		faceDown: Deck{dealerCards[len(dealerCards)-1]},
		total:    0,
	}
	dealer.total = getTotal(dealer.faceUp)

	player := Player{
		faceUp:   playerCards,
		faceDown: nil,
		total:    0,
	}
	player.total = getTotal(player.faceUp)

	fmt.Println("\nHere are my cards: ")
	dealer.print()

	fmt.Println("\nHere are your cards: ")
	player.print()

	if player.total == 21 {
		fmt.Println("BLACKJACK!")

		dealer.faceUp = append(dealer.faceUp, dealer.faceDown...)
		dealer.faceDown = nil
		dealer.total = getTotal(dealer.faceUp)

		fmt.Println("\nMy cards are: ")
		dealer.print()

		if dealer.total == 21 {
			fmt.Println("\nDealer has BLACKJACK! It's a TIE.")
		} else {
			fmt.Println("\nYou WIN")
		}
	} else {
		getUserChoice()
	}
}

func getUserChoice() {
	var choice string

	fmt.Print("\nDo you want to hit or stand? (h/s): ")
	fmt.Scanln(&choice)
	choice = strings.ToLower(strings.TrimSpace(choice))

	if choice != "h" && choice != "s" {
		fmt.Println("Invalid choice. Please enter 'h' for hit or 's' for stand.")
		getUserChoice()
	} else if choice == "h" {
		fmt.Println("Hitting!")
	} else if choice == "s" {
		fmt.Println("Standing.")
	}
}

func (p Player) handleHit() {

}

func (p Player) handleStand() {

}

func getTotal(cards Deck) int {
	total := 0
	acesIndex := []int{}

	for index, card := range cards {
		if card.value != Ace {
			total += card.num[0]
		} else {
			acesIndex = append(acesIndex, index)
		}
	}

	for i := 0; i < len(acesIndex); i++ {
		lowerValue, higherValue := cards[acesIndex[i]].num[0], cards[acesIndex[i]].num[1]
		if total+higherValue > 21 {
			total += lowerValue
		} else {
			total += higherValue
		}
	}

	return total
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

func (p Player) print() {
	fmt.Println()

	if p.faceUp != nil {
		for _, card := range p.faceUp {
			printCard(card, true)
		}
	}

	if p.faceDown != nil {
		for _, card := range p.faceDown {
			printCard(card, false)
		}
	}

	fmt.Println("Total: ", p.total)
}

func printCard(card Card, show bool) {
	width := 9
	height := 7

	fmt.Println("┌─" + strings.Repeat("─", width-2) + "─┐")

	if show {
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
