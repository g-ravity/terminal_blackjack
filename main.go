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

type PlayerStatus string

const (
	Hitting  PlayerStatus = "h"
	Standing PlayerStatus = "s"
)

type Player struct {
	faceUp   Deck
	faceDown Deck
	total    int
	status   PlayerStatus
}

var dealer = Player{
	faceUp:   nil,
	faceDown: nil,
	total:    0,
	status:   Hitting,
}
var player = Player{
	faceUp:   nil,
	faceDown: nil,
	total:    0,
	status:   Hitting,
}

var cardsDeck Deck

func printWithTypingEffect(text string, delay int) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println()
	time.Sleep(time.Duration(delay) * time.Millisecond)
}

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

func getPlayerChoice() {
	var choice string

	printWithTypingEffect("\nDo you want to hit or stand? (h/s): ", 300)
	fmt.Scanln(&choice)
	choice = strings.ToLower(strings.TrimSpace(choice))

	if choice != "h" && choice != "s" {
		printWithTypingEffect("Invalid choice. Please enter 'h' for hit or 's' for stand.", 300)
		getPlayerChoice()
	} else if choice == "h" {
		printWithTypingEffect("Hitting!", 300)

		player.handleHit()
	} else if choice == "s" {
		printWithTypingEffect("Standing.", 300)

		player.handleStand()
	}
}

func getDealerChoice() {
	if dealer.total < 17 {
		printWithTypingEffect("\nDealer is Hitting!", 500)

		dealer.handleHit()
	} else {
		printWithTypingEffect("\nDealer is Standing.", 500)

		dealer.handleStand()
	}
}

func revealDealerCards() {
	dealer.faceUp = append(dealer.faceUp, dealer.faceDown...)
	dealer.faceDown = nil
	dealer.total = getTotal(dealer.faceUp)

	printWithTypingEffect("\nMy cards are: ", 1000)
	dealer.print()
}

func isBlackJack(p Player) bool {
	return len(p.faceUp) == 2 && p.total == 21
}

func checkGameLogic() bool {
	isGameOver := false

	if player.total > 21 {
		printWithTypingEffect("BUST!", 500)
		printWithTypingEffect("\nYou LOST", 500)

		isGameOver = true
	} else if player.status == Standing {
		if isBlackJack(player) {
			printWithTypingEffect("\nBLACKJACK!", 500)
		}

		if len(dealer.faceDown) > 0 {
			revealDealerCards()
		}

		if isBlackJack(player) {
			if isBlackJack(dealer) {
				printWithTypingEffect("\nDealer has BLACKJACK! It's a TIE.", 500)
			} else {
				printWithTypingEffect("\nYou WIN", 500)
			}

			isGameOver = true
		} else if isBlackJack(dealer) {
			printWithTypingEffect("\nDealer has BLACKJACK! You LOST.", 500)

			isGameOver = true
		} else if dealer.total > 21 {
			printWithTypingEffect("\nDealer BUST! You WIN.", 500)

			isGameOver = true
		} else if dealer.status == Standing {
			if player.total < dealer.total {
				printWithTypingEffect("\nDealer has a higher total. You LOST.", 500)
			} else if player.total > dealer.total {
				printWithTypingEffect("\nYou have a higher total. You WIN.", 500)
			} else {
				printWithTypingEffect("\nIt's a TIE.", 500)
			}

			isGameOver = true
		}
	}

	return isGameOver
}

func getChoice() {
	if player.status == Hitting {
		getPlayerChoice()
	} else if player.status == Standing {
		getDealerChoice()
	}
}

func (p *Player) handleHit() {
	var playerCard Deck

	playerCard, cardsDeck = cardsDeck.deal(1)
	p.addToHand(playerCard)
	p.print()

	if p.total == 21 {
		p.handleStand()

		return
	}

	if !checkGameLogic() {
		getChoice()
	}
}

func (p *Player) handleStand() {
	p.status = Standing

	if !checkGameLogic() {
		getChoice()
	}
}

func (p *Player) addToHand(cards Deck) {
	p.faceUp = append(p.faceUp, cards...)
	p.total = getTotal(p.faceUp)
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
