package main

func isBlackJack(p Player) bool {
	return len(p.faceUp) == 2 && p.total == 21
}

func checkForFinish() bool {
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

func (p *Player) handleHit() {
	var playerCard Deck

	playerCard, cardsDeck = cardsDeck.deal(1)
	p.addToHand(playerCard)
	p.print()

	if p.total == 21 {
		p.handleStand()

		return
	}

	if !checkForFinish() {
		getChoice()
	}
}

func (p *Player) handleStand() {
	p.status = Standing

	if !checkForFinish() {
		getChoice()
	}
}
