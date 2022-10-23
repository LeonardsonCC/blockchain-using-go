package main

import (
	"time"

	"github.com/LeonardsonCC/blockchain-using-go/blockchain"
	"github.com/LeonardsonCC/blockchain-using-go/ui"
)

func main() {
	moment := time.Date(2022, 10, 22, 20, 12, 0, 0, time.Now().Location())
	bc := blockchain.NewBlockchain(moment, 2)

	ui.Start(bc)
}
