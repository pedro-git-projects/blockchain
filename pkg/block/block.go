package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	nonce        int
	previousHash [32]byte
	timeStamp    int64
	transactions []string
}

// MarshalJSON overloads our json Marshll for blocks exposing  its private struct fields for our hash method  as well as marshaling them into JSON
func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Nonce        int      `json:"nonce"`
		PreviousHash [32]byte `json:"previous_hash"`
		Timestamp    int64    `json:"timestamp"`
		Transactions []string `json:"transactions"`
	}{
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Timestamp:    b.timeStamp,
		Transactions: b.transactions,
	})
}

// Hash is a method that takes the JSON marshal of a block and hashes it using a sh256 cryptographic function
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	//	fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

// NewBlock returns a pointer to a new block containing the current timestamp
func NewBlock(nonce int, previousHash [32]byte) *Block {
	b := new(Block)
	b.timeStamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}

// Print is a method to improve our block readability
func (b Block) Print() {
	fmt.Printf("timestamp %d\n", b.timeStamp)
	fmt.Printf("previous_hash %x\n", b.previousHash)
	fmt.Printf("nonce %d\n", b.nonce)
	fmt.Printf("transactions %v\n", b.transactions)
}
