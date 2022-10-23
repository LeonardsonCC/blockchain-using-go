package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"
)

type Block struct {
	Timestamp     time.Time `json:"timestamp"`
	Txs           []Tx      `json:"txs"`
	Hash          Hash      `json:"hash"`
	PrecedingHash Hash      `json:"preceding_block"`
	Difficulty    int       `json:"difficulty"`
}

func NewBlock(timestamp time.Time, txs []Tx, precedingHash Hash) *Block {
	bl := &Block{
		Timestamp:     timestamp,
		Txs:           txs,
		PrecedingHash: precedingHash,
		Difficulty:    2,
	}

	bl.Hash = bl.computeHash()

	return bl
}

func (b *Block) AddTx(timestamp time.Time, data *Data) {
	// getting the last tx hash
	tx := NewTx(timestamp, *data, b.Txs[len(b.Txs)-1].Hash, b.Difficulty)
	b.Txs = append(b.Txs, *tx)
	b.Hash = b.computeHash()
}

func (t *Block) computeHash() Hash {
	toHash := map[string]interface{}{
		"timestamp":      t.Timestamp,
		"txs":            t.Txs,
		"preceding_hash": t.PrecedingHash,
	}
	hash, _ := json.Marshal(toHash)

	hasher := sha256.New()
	hasher.Write(hash)
	h := hex.EncodeToString(hasher.Sum(nil))

	return Hash(h)
}

func (t *Block) ToString() string {
	s, _ := json.Marshal(t)
	return string(s)
}
