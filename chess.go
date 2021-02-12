package main 

import (
	"math/rand"
	"time"
)

type Chess struct {
	Board [][]string
	Players []Player
}

func (c *Chess) Init() {
	board := make([][]string, 8, 8)
	c.Board = board

	
}

func (d *Chess) Move(*) {
	
}