package main

import (
	"fmt"
	"strings"
)

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
	if dealer.total < 17 || (dealer.total == 17 && containsValue(dealer.faceUp, Ace)) {
		printWithTypingEffect("\nDealer is Hitting!", 500)

		dealer.handleHit()
	} else {
		printWithTypingEffect("\nDealer is Standing.", 500)

		dealer.handleStand()
	}
}

func getChoice() {
	if player.status == Hitting {
		getPlayerChoice()
	} else if player.status == Standing {
		getDealerChoice()
	}
}

func revealDealerCards() {
	dealer.faceUp = append(dealer.faceUp, dealer.faceDown...)
	dealer.faceDown = nil
	dealer.total = getTotal(dealer.faceUp)

	printWithTypingEffect("\nMy cards are: ", 1000)
	dealer.print()
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
		lowerValue := cards[acesIndex[i]].num[0]
		total += lowerValue
	}

	if len(acesIndex) > 0 {
		aceValueDiff := cards[acesIndex[0]].num[1] - cards[acesIndex[0]].num[0]

		if 21-total >= aceValueDiff {
			total += aceValueDiff
		}
	}

	return total
}
