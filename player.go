package main

import (
	"fmt"
)

type Player struct {
	Name string
	Age string
	Wallet WalletDto
}

type WalletDto struct {
	Amount float64
}

func (u *Player) InitNewPlayer() {

}

func (u *Player) Bet (amount float64) error {
	walletData := (*u).Wallet

	if walletData.Amount - amount < 0 {
		return fmt.Errorf("Insufficient balance")
	
	}

	walletData.Amount = walletData.Amount - amount
	u.Wallet = walletData
	return nil
}
