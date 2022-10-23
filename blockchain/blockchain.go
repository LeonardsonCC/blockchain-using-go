package blockchain

import (
	"encoding/json"
	"time"
)

type Blockchain struct {
	Timestamp   time.Time `json:"timestamp"`
	Blocks      []Block   `json:"blocks"`
	Hash        Hash      `json:"hash"`
	TxsPerBlock int       `json:"txs_per_block"`
}

func NewBlockchain(timestamp time.Time, difficulty int) *Blockchain {
	bc := &Blockchain{
		Timestamp: timestamp,
		Blocks: []Block{
			*NewBlock(timestamp, []Tx{*NewTx(timestamp, *NewData("null", "null", 0), "0", difficulty)}, "0"), // genesis block
		},
		TxsPerBlock: 10, // default value for testing
	}

	return bc
}

func (b *Blockchain) AddTx(timestamp time.Time, data *Data) {
	lastBlock := b.Blocks[len(b.Blocks)-1]
	if len(lastBlock.Txs) < b.TxsPerBlock {
		lastBlock.AddTx(timestamp, data)
		b.Blocks[len(b.Blocks)-1] = lastBlock
	} else {
		newBlock := NewBlock(time.Now(), []Tx{}, lastBlock.Hash)
		newBlock.AddTx(time.Now(), data)
		b.Blocks = append(b.Blocks, *newBlock)
	}
}

func (b *Blockchain) ToString() string {
	s, _ := json.Marshal(b)
	return string(s)
}
