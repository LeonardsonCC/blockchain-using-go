package main

import (
	"log"
	"time"

	"github.com/LeonardsonCC/blockchain-using-go/block"
)

func main() {
	t := time.Now()

	moment := time.Date(2022, 10, 22, 20, 12, 0, 0, time.Now().Location())

	bc := block.NewBlockchain(moment, 5)

	bc.AddTx(moment, block.NewData("TEST-1", "TEST-2", 100))
	bc.AddTx(moment, block.NewData("TEST-2", "TEST-3", 200))
	bc.AddTx(moment, block.NewData("TEST-3", "TEST-1", 300))

	log.Printf("Blockchain: %s\n", bc.ToString())

	log.Println(time.Since(t))
}
