package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var cardsDeck Deck

func main() {
	printWithTypingEffect("Hello, Welcome to the game of Blackjack!", 1000)
	printWithTypingEffect("I am the dealer.", 1000)

	cardsDeck = createDeck()
	cardsDeck.shuffle()

	numToDeal := 2

	dealerCards, cardsDeck := cardsDeck.deal(numToDeal)
	playerCards, cardsDeck := cardsDeck.deal(numToDeal)

	dealer.faceUp = dealerCards[:len(dealerCards)-1]
	dealer.faceDown = Deck{dealerCards[len(dealerCards)-1]}
	dealer.total = getTotal(dealer.faceUp)

	player.faceUp = playerCards
	player.total = getTotal(player.faceUp)

	printWithTypingEffect("\nHere are my cards: ", 1000)
	dealer.print()

	printWithTypingEffect("\nHere are your cards: ", 1000)
	player.print()

	time.Sleep(300 * time.Millisecond)

	if player.total == 21 {
		player.handleStand()
	} else {
		getChoice()
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

func (p Player) print() {
	fmt.Println()

	if p.faceUp != nil {
		for _, card := range p.faceUp {
			printCard(card, true)
			time.Sleep(500 * time.Millisecond)
		}
	}

	if p.faceDown != nil {
		for _, card := range p.faceDown {
			printCard(card, false)
			time.Sleep(500 * time.Millisecond)
		}
	}

	printWithTypingEffect("Total: ", 300)
	time.Sleep(200 * time.Millisecond)
	fmt.Println(p.total)
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
