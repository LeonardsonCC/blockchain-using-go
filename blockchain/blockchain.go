package blockchain

import (
	"encoding/json"
	"time"
)

type Blockchain struct {
	Timestamp time.Time `json:"timestamp"`
	Blocks    []Block   `json:"blocks"`
	Hash      Hash      `json:"hash"`
}

func NewBlockchain(timestamp time.Time, difficulty int) *Blockchain {
	bc := &Blockchain{
		Timestamp: timestamp,
		Blocks: []Block{
			*NewBlock(timestamp, []Tx{*NewTx(timestamp, *NewData("null", "null", 0), "0", difficulty)}, "0"), // genesis block
		},
	}

	return bc
}

func (b *Blockchain) AddTx(timestamp time.Time, data *Data) {
	lastBlock := b.Blocks[len(b.Blocks)-1]
	lastBlock.AddTx(timestamp, data)

	// updates the block, because it's not a pointer
	b.Blocks[len(b.Blocks)-1] = lastBlock
}

func (b *Blockchain) ToString() string {
	s, _ := json.Marshal(b)
	return string(s)
}
