package main

import (
	"testing"
)

func isUnique(deck Deck) bool {
	seen := make(map[string]bool)
	for _, card := range deck {
		if seen[string(card.suit)+string(card.value)] {
			return false
		}
		seen[string(card.suit)+string(card.value)] = true
	}
	return true
}

func TestCreateDeck(t *testing.T) {
	deck := createDeck()

	if len(deck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(deck))
	}

	if !isUnique(deck) {
		t.Errorf("Deck contains duplicate cards")
	}
}

func TestShuffleDeck(t *testing.T) {
	deck := createDeck()

	originalDeck := make(Deck, len(deck))
	copy(originalDeck, deck)

	deck.shuffle()

	if len(deck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(deck))
	}

	if !isUnique(deck) {
		t.Errorf("Shuffled Deck contains duplicate cards")
	}

	isShuffled := false

	for i, card := range originalDeck {
		if card.suit != deck[i].suit || card.value != deck[i].value {
			isShuffled = true
			break
		}
	}

	if !isShuffled {
		t.Errorf("Shuffled deck is same as original deck")
	}
}

func TestDeckDeal(t *testing.T) {
	deck := createDeck()
	hand, deck := deck.deal(5)

	if len(hand) != 5 {
		t.Errorf("Expected hand length of 5, but got %v", len(hand))
	}

	if len(deck) != 47 {
		t.Errorf("Expected remaining deck length of 47, but got %v", len(deck))
	}

	if !isUnique(hand) {
		t.Errorf("Hand contains duplicate cards")
	}

	if !isUnique(deck) {
		t.Errorf("Remaining deck contains duplicate cards")
	}

	for _, card := range hand {
		if containsCard(deck, card) {
			t.Errorf("Remaining deck contains card from hand")
			break
		}
	}
}

func TestHandTotal(t *testing.T) {
	handCases := map[int]struct {
		hand  Deck
		total int
	}{
		0: {
			hand:  Deck{Card{suit: Hearts, value: Ace}, Card{suit: Spades, value: Two}, Card{suit: Diamonds, value: Ace}, Card{suit: Clubs, value: Seven}},
			total: 21,
		},
		1: {
			hand:  Deck{Card{suit: Hearts, value: Ace}, Card{suit: Spades, value: Ace}, Card{suit: Diamonds, value: Ace}, Card{suit: Clubs, value: Eight}},
			total: 21,
		},
		2: {
			hand:  Deck{Card{suit: Hearts, value: Ace}, Card{suit: Hearts, value: Ace}, Card{suit: Spades, value: Ace}, Card{suit: Diamonds, value: Ace}, Card{suit: Clubs, value: Ace}, Card{suit: Hearts, value: Ace}, Card{suit: Spades, value: Six}},
			total: 12,
		},
		3: {
			hand:  Deck{Card{suit: Hearts, value: Ace}, Card{suit: Spades, value: Five}, Card{suit: Diamonds, value: Five}, Card{suit: Clubs, value: Ace}},
			total: 12,
		},
		4: {
			hand:  Deck{Card{suit: Hearts, value: Ace}, Card{suit: Spades, value: Ace}, Card{suit: Diamonds, value: Nine}, Card{suit: Clubs, value: Ace}},
			total: 12,
		},
		5: {
			hand:  Deck{Card{suit: Hearts, value: Three}, Card{suit: Spades, value: Six}, Card{suit: Diamonds, value: Ace}, Card{suit: Clubs, value: Ace}},
			total: 21,
		},
		6: {
			hand:  Deck{Card{suit: Hearts, value: Ten}, Card{suit: Spades, value: Ace}, Card{suit: Diamonds, value: Ten}},
			total: 21,
		},
		7: {
			hand:  Deck{Card{suit: Hearts, value: Five}, Card{suit: Spades, value: Five}, Card{suit: Diamonds, value: Five}, Card{suit: Clubs, value: Five}, Card{suit: Hearts, value: Ace}},
			total: 21,
		},
		8: {
			hand:  Deck{Card{suit: Hearts, value: Nine}, Card{suit: Spades, value: Ace}, Card{suit: Diamonds, value: Ace}, Card{suit: Clubs, value: Ace}},
			total: 12,
		},
		9: {
			hand:  Deck{Card{suit: Hearts, value: Ten}, Card{suit: Spades, value: Nine}, Card{suit: Diamonds, value: Ace}, Card{suit: Clubs, value: Ace}},
			total: 21,
		},
		10: {
			hand:  Deck{Card{suit: Hearts, value: King}, Card{suit: Spades, value: Queen}},
			total: 20,
		},
		11: {
			hand:  Deck{Card{suit: Hearts, value: Joker}, Card{suit: Spades, value: King}, Card{suit: Diamonds, value: Ace}},
			total: 21,
		},
		12: {
			hand:  Deck{Card{suit: Hearts, value: Queen}, Card{suit: Spades, value: Joker}, Card{suit: Diamonds, value: King}},
			total: 30,
		},
		13: {
			hand:  Deck{Card{suit: Hearts, value: King}, Card{suit: Spades, value: Ace}},
			total: 21,
		},
		14: {
			hand:  Deck{Card{suit: Hearts, value: Ace}, Card{suit: Spades, value: Ace}},
			total: 12,
		},
		15: {
			hand:  Deck{Card{suit: Hearts, value: Ace}, Card{suit: Spades, value: Ace}, Card{suit: Diamonds, value: Ace}},
			total: 13,
		},
		16: {
			hand:  Deck{Card{suit: Hearts, value: Ace}, Card{suit: Spades, value: Ace}, Card{suit: Diamonds, value: Ace}, Card{suit: Clubs, value: Ace}, Card{suit: Hearts, value: Seven}},
			total: 21,
		},
		17: {
			hand:  Deck{Card{suit: Hearts, value: Ace}, Card{suit: Spades, value: Ace}, Card{suit: Diamonds, value: Five}, Card{suit: Clubs, value: Five}},
			total: 12,
		},
		18: {
			hand:  Deck{Card{suit: Hearts, value: Ace}, Card{suit: Spades, value: Six}},
			total: 17,
		},
		19: {
			hand:  Deck{Card{suit: Hearts, value: Ace}, Card{suit: Spades, value: Nine}},
			total: 20,
		},
		20: {
			hand:  Deck{Card{suit: Hearts, value: Ace}, Card{suit: Spades, value: Ten}},
			total: 21,
		},
		21: {
			hand:  Deck{Card{suit: Hearts, value: Ace}, Card{suit: Spades, value: Eight}, Card{suit: Diamonds, value: Two}},
			total: 21,
		},
		22: {
			hand:  Deck{Card{suit: Hearts, value: Ace}, Card{suit: Spades, value: Six}, Card{suit: Diamonds, value: Six}},
			total: 13,
		},
		23: {
			hand:  Deck{Card{suit: Hearts, value: Ace}, Card{suit: Spades, value: Ace}, Card{suit: Diamonds, value: Nine}},
			total: 21,
		},
		24: {
			hand:  Deck{Card{suit: Hearts, value: Two}, Card{suit: Spades, value: Three}},
			total: 5,
		},
		25: {
			hand:  Deck{Card{suit: Hearts, value: Ten}, Card{suit: Spades, value: Ten}},
			total: 20,
		},
		26: {
			hand:  Deck{Card{suit: Hearts, value: Ten}, Card{suit: Spades, value: Nine}, Card{suit: Diamonds, value: Two}},
			total: 21,
		},
		27: {
			hand:  Deck{Card{suit: Hearts, value: Ten}, Card{suit: Spades, value: Ten}, Card{suit: Diamonds, value: Ten}},
			total: 30,
		},
	}

	for i, tc := range handCases {
		for j := range tc.hand {
			tc.hand[j].num = (tc.hand[j].value).getNum()
		}

		total := getTotal(tc.hand)
		if total != tc.total {
			t.Errorf("Test case %d: Expected total of %d, but got %d", i, tc.total, total)
		}
	}
}
