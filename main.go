package main

import (
	"fmt"
)

func main () {
	var deck Deck
	var user User

	deck.Init()
	user.InitNewUser()

	deck.Shuffle()
	fmt.Println(deck)

	user.Bet(5.0)
}
