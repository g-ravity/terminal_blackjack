package main

import (
	"fmt"
	"time"
)

func printWithTypingEffect(text string, delay int) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println()
	time.Sleep(time.Duration(delay) * time.Millisecond)
}

func containsValue(cards Deck, value Value) bool {
	for _, card := range cards {
		if card.value == value {
			return true
		}
	}

	return false
}

func containsSuit(cards Deck, suit Suit) bool {
	for _, card := range cards {
		if card.suit == suit {
			return true
		}
	}
	return false
}

func containsCard(cards Deck, card Card) bool {
	for _, c := range cards {
		if c.suit == card.suit && c.value == card.value {
			return true
		}
	}
	return false
}
