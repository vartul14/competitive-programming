package main 

import (
	"math/rand"
	"time"
)

type Card struct {
	Type string
	Suit string
}

type Deck []Card

func (d *Deck) Init() {


	types := []string {"one", "two"}
	suits := []string {"spade", "heart"}

	for i := 0; i < len(types); i++ {
		for j := 0; j < len(suits); j++ {
			c := Card{
				Type: types[i],
				Suit: suits[j],
			}
			*d = append(*d, c)
		}
	}
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(*d), func(i, j int) {(*d)[i], (*d)[j] = (*d)[j], (*d)[i]})
}