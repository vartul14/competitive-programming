package main

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	var deck Deck
	deck.Init()
	//t.Errorf("error in deck init")
	fmt.Println(deck[0])
}