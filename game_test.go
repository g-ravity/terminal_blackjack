package main

import (
	"testing"
)

func TestBlackjack(t *testing.T) {
	tests := []struct {
		name        string
		player      Player
		isBlackjack bool
	}{
		{
			name: "21 with Ace and Ten",
			player: Player{
				faceUp: Deck{
					Card{suit: Hearts, value: Ace, num: Ace.getNum()},
					Card{suit: Spades, value: Ten, num: Ten.getNum()},
				},
			},
			isBlackjack: true,
		},
		{
			name: "21 without Ace and Ten",
			player: Player{
				faceUp: Deck{
					Card{suit: Hearts, value: Ace, num: Ace.getNum()},
					Card{suit: Spades, value: Six, num: Six.getNum()},
					Card{suit: Clubs, value: Four, num: Four.getNum()},
				},
			},
			isBlackjack: false,
		},
		{
			name: "Not 21",
			player: Player{
				faceUp: Deck{
					Card{suit: Hearts, value: King, num: King.getNum()},
					Card{suit: Spades, value: Seven, num: Seven.getNum()},
				},
			},
			isBlackjack: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.player.total = getTotal(tc.player.faceUp)

			if !tc.isBlackjack && isBlackJack(tc.player) {
				t.Errorf("Expected %v to not be blackjack", tc.player)
			} else if tc.isBlackjack && !isBlackJack(tc.player) {
				t.Errorf("Expected %v to be blackjack", tc.player)
			}
		})
	}

	resetGlobals()
}

func TestCheckForFinish(t *testing.T) {
	tests := []struct {
		name     string
		setup    func()
		wantOver bool
	}{
		{
			name: "Player busts",
			setup: func() {
				player = Player{
					faceUp: Deck{
						Card{suit: Hearts, value: Ten},
						Card{suit: Spades, value: Ten},
						Card{suit: Diamonds, value: Five},
					},
					total:  25,
					status: Hitting,
				}
			},
			wantOver: true,
		},
		{
			name: "Player blackjack, dealer no blackjack",
			setup: func() {
				player = Player{
					faceUp: Deck{
						Card{suit: Hearts, value: Ace},
						Card{suit: Spades, value: Ten},
					},
					total:  21,
					status: Standing,
				}
				dealer = Player{
					faceUp: Deck{
						Card{suit: Hearts, value: Ten},
						Card{suit: Spades, value: Nine},
					},
					total:  19,
					status: Standing,
				}
			},
			wantOver: true,
		},
		{
			name: "Both have blackjack",
			setup: func() {
				player = Player{
					faceUp: Deck{
						Card{suit: Hearts, value: Ace},
						Card{suit: Spades, value: King},
					},
					total:  21,
					status: Standing,
				}
				dealer = Player{
					faceUp: Deck{
						Card{suit: Diamonds, value: Queen},
						Card{suit: Clubs, value: Ace},
					},
					total:  21,
					status: Standing,
				}
			},
			wantOver: true,
		},
		{
			name: "Dealer blackjack, player no blackjack",
			setup: func() {
				player = Player{
					faceUp: Deck{
						Card{suit: Hearts, value: Ten},
						Card{suit: Spades, value: Nine},
					},
					total:  19,
					status: Standing,
				}
				dealer = Player{
					faceUp: Deck{
						Card{suit: Diamonds, value: Ace},
						Card{suit: Clubs, value: Ten},
					},
					total:  21,
					status: Standing,
				}
			},
			wantOver: true,
		},
		{
			name: "Dealer busts",
			setup: func() {
				player = Player{
					faceUp: Deck{
						Card{suit: Hearts, value: Ten},
						Card{suit: Spades, value: Eight},
					},
					total:  18,
					status: Standing,
				}
				dealer = Player{
					faceUp: Deck{
						Card{suit: Diamonds, value: Ten},
						Card{suit: Clubs, value: Ten},
						Card{suit: Hearts, value: Five},
					},
					total:  25,
					status: Standing,
				}
			},
			wantOver: true,
		},
		{
			name: "Dealer wins with higher total",
			setup: func() {
				player = Player{
					faceUp: Deck{
						Card{suit: Hearts, value: Ten},
						Card{suit: Spades, value: Seven},
					},
					total:  17,
					status: Standing,
				}
				dealer = Player{
					faceUp: Deck{
						Card{suit: Diamonds, value: Ten},
						Card{suit: Clubs, value: Nine},
					},
					total:  19,
					status: Standing,
				}
			},
			wantOver: true,
		},
		{
			name: "Player wins with higher total",
			setup: func() {
				player = Player{
					faceUp: Deck{
						Card{suit: Hearts, value: Ten},
						Card{suit: Spades, value: Nine},
					},
					total:  19,
					status: Standing,
				}
				dealer = Player{
					faceUp: Deck{
						Card{suit: Diamonds, value: Ten},
						Card{suit: Clubs, value: Seven},
					},
					total:  17,
					status: Standing,
				}
			},
			wantOver: true,
		},
		{
			name: "Tie game",
			setup: func() {
				player = Player{
					faceUp: Deck{
						Card{suit: Hearts, value: Ten},
						Card{suit: Spades, value: Eight},
					},
					total:  18,
					status: Standing,
				}
				dealer = Player{
					faceUp: Deck{
						Card{suit: Diamonds, value: Ten},
						Card{suit: Clubs, value: Eight},
					},
					total:  18,
					status: Standing,
				}
			},
			wantOver: true,
		},
		{
			name: "Game continues - player hitting",
			setup: func() {
				player = Player{
					faceUp: Deck{
						Card{suit: Hearts, value: Ten},
						Card{suit: Spades, value: Five},
					},
					total:  15,
					status: Hitting,
				}
			},
			wantOver: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			result := checkForFinish()

			if result != tt.wantOver {
				t.Errorf("checkForFinish() = %v, want %v", result, tt.wantOver)
			}
		})
	}

	resetGlobals()
}
