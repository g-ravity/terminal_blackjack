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
			tt.setup()
			for i, card := range dealer.faceUp {
				dealer.faceUp[i].num = card.value.getNum()
			}
			dealer.total = getTotal(dealer.faceUp)

			mock := &mockPlayer{hitCalled: false, standCalled: false}

			getDealerChoice(mock)

			if tt.wantHit && !mock.hitCalled {
				t.Error("Expected dealer to hit but they didn't")
			}
			if !tt.wantHit && !mock.standCalled {
				t.Error("Expected dealer to stand but they didn't")
			}
		})
	}
}
