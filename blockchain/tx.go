package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strings"
	"time"
)

type Data struct {
	Sender   Address `json:"sender"`
	Receiver Address `json:"receiver"`
	Value    int     `json:"value"`
}

type Tx struct {
	Timestamp     time.Time `json:"timestamp"`
	Data          Data      `json:"data"`
	PrecedingHash Hash      `json:"preceding_hash"`
	Hash          Hash      `json:"hash"`
	Nonce         int       `json:"nonce"`
}

func NewTx(timestamp time.Time, data Data, precedingHash Hash, difficulty int) *Tx {
	tx := &Tx{
		Timestamp:     timestamp,
		Data:          data,
		PrecedingHash: precedingHash,
	}

	tx.Hash = tx.computeHash(difficulty)

	return tx
}

func NewData(sender, receiver string, value int) *Data {
	return &Data{
		Sender:   Address(sender),
		Receiver: Address(receiver),
		Value:    value,
	}
}

func (t *Tx) computeHash(difficulty int) Hash {
	nonce := 0
	for {
		data, err := json.Marshal(t.Data)
		if err != nil {
			return ""
		}

		toHash := map[string]interface{}{
			"timestamp":      t.Timestamp,
			"preceding_hash": t.PrecedingHash,
			"data":           data,
			"nonce":          nonce,
		}
		hash, err := json.Marshal(toHash)

		hasher := sha256.New()
		hasher.Write(hash)
		h := hex.EncodeToString(hasher.Sum(nil))

		if h[0:difficulty] == strings.Repeat("0", difficulty) {
			return Hash(h)
		}
		nonce = nonce + 1
	}
}

func (t *Tx) ToString() string {
	s, _ := json.Marshal(t)
	return string(s)
}
