package main

import (
	"testing"
)

type mockPlayer struct {
	hitCalled   bool
	standCalled bool
}

func (m *mockPlayer) handleHit()   { m.hitCalled = true }
func (m *mockPlayer) handleStand() { m.standCalled = true }

func TestGetDealerChoice(t *testing.T) {
	tests := []struct {
		name    string
		setup   func()
		wantHit bool
	}{
		{
			name: "Dealer hits when total < 17",
			setup: func() {
				dealer = Player{
					faceUp: Deck{
						Card{suit: Hearts, value: Six},
						Card{suit: Spades, value: Ten},
					},
					status: Hitting,
				}
			},
			wantHit: true,
		},
		{
			name: "Dealer hits when total = 17 with Ace with 2 cards",
			setup: func() {
				dealer = Player{
					faceUp: Deck{
						Card{suit: Hearts, value: Ace},
						Card{suit: Spades, value: Six},
					},
					status: Hitting,
				}
			},
			wantHit: true,
		},
		{
			name: "Dealer hits when total = 17 with Ace with more than 2 cards",
			setup: func() {
				dealer = Player{
					faceUp: Deck{
						Card{suit: Hearts, value: Ace},
						Card{suit: Spades, value: Two},
						Card{suit: Spades, value: Four},
					},
					status: Hitting,
				}
			},
			wantHit: true,
		},
		{
			name: "Dealer stands when total = 17 without Ace",
			setup: func() {
				dealer = Player{
					faceUp: Deck{
						Card{suit: Hearts, value: Ten},
						Card{suit: Spades, value: Seven},
					},
					status: Hitting,
				}
			},
			wantHit: false,
		},
		{
			name: "Dealer stands when total > 17 with Ace",
			setup: func() {
				dealer = Player{
					faceUp: Deck{
						Card{suit: Hearts, value: Ace},
						Card{suit: Spades, value: Nine},
					},
					status: Hitting,
				}
			},
			wantHit: false,
		},
		{
			name: "Dealer stands when total > 17 without Ace",
			setup: func() {
				dealer = Player{
					faceUp: Deck{
						Card{suit: Hearts, value: Ten},
						Card{suit: Spades, value: Nine},
					},
					status: Hitting,
				}
			},
			wantHit: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := mockPlayer{hitCalled: false, standCalled: false}
			tt.setup()

			for i, card := range dealer.faceUp {
				dealer.faceUp[i].num = card.value.getNum()
			}
			dealer.total = getTotal(dealer.faceUp)

			getDealerChoice(&mock)

			if tt.wantHit && !mock.hitCalled {
				t.Error("Expected dealer to hit but they didn't")
			}
			if !tt.wantHit && !mock.standCalled {
				t.Error("Expected dealer to stand but they didn't")
			}
		})
	}

	resetGlobals()
}

func TestAddToHand(t *testing.T) {
	tests := []struct {
		name        string
		initialHand Deck
		cardsToAdd  Deck
		wantTotal   int
		wantFaceUp  int
	}{
		{
			name:        "Add single card to empty hand",
			initialHand: nil,
			cardsToAdd: Deck{
				Card{suit: Hearts, value: Ten, num: Ten.getNum()},
			},
			wantTotal:  10,
			wantFaceUp: 1,
		},
		{
			name:        "Add multiple cards to empty hand",
			initialHand: nil,
			cardsToAdd: Deck{
				Card{suit: Hearts, value: Ten, num: Ten.getNum()},
				Card{suit: Spades, value: Five, num: Five.getNum()},
			},
			wantTotal:  15,
			wantFaceUp: 2,
		},
		{
			name: "Add card to existing hand",
			initialHand: Deck{
				Card{suit: Hearts, value: Ten, num: Ten.getNum()},
			},
			cardsToAdd: Deck{
				Card{suit: Spades, value: Five, num: Five.getNum()},
			},
			wantTotal:  15,
			wantFaceUp: 2,
		},
		{
			name: "Add Ace to existing hand",
			initialHand: Deck{
				Card{suit: Hearts, value: Six, num: Six.getNum()},
			},
			cardsToAdd: Deck{
				Card{suit: Spades, value: Ace, num: Ace.getNum()},
			},
			wantTotal:  17,
			wantFaceUp: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Player{
				faceUp: tt.initialHand,
				total:  getTotal(tt.initialHand),
			}

			p.addToHand(tt.cardsToAdd)

			if p.total != tt.wantTotal {
				t.Errorf("addToHand() got total = %v, want %v", p.total, tt.wantTotal)
			}

			if len(p.faceUp) != tt.wantFaceUp {
				t.Errorf("addToHand() got faceUp length = %v, want %v", len(p.faceUp), tt.wantFaceUp)
			}
		})
	}

	resetGlobals()
}
