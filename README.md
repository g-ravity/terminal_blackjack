# Terminal Blackjack

A simple stripped down version of the Blackjack game for the terminal enthusiasts.

## Usage

You can clone the repository and run the game with the following command:

```
go run main.go
```

OR

You can download the executable file called `terminal_blackjack` (download the .exe file for windows) and run it directly.

## Gameplay

Ace has a value of 1 / 11 (whichever is more favorable). Face Cards (K/Q/J) have a value of 10. Rest of the cards have the same value as their rank.

The game will start with the dealer showing his cards. His `hole card` will be hidden.

Then you'll be dealt your cards. You can either `hit` or `stand`.

If you get a `blackjack` (Ace + K/Q/J) on the first hand, your hand ends and the dealer will reveal his hole card.

The dealer has to hit if his total is below 17, otherwise he stands.

If you go over 21, you bust and lose. If the dealer busts, you win. If both you and the dealer have the same total, it's a tie.

`blackjack` beats any other kind of 21. So even if you have 21 but not a `blackjack`, and the dealer gets a `blackjack`, you lose. Same thing goes other way around.

## Future

I will be adding more features to the game, like betting, statistics, multiplayer mode etc. So star the repository and watch for updates.
