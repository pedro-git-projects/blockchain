package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	timeStamp    int64
	nonce        int
	previousHash [32]byte
	transactions []*Transaction
}

// MarshalJSON overloads our json Marshll for blocks exposing  its private struct fields for our hash method  as well as marshaling them into JSON
func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Nonce        int            `json:"nonce"`
		PreviousHash [32]byte       `json:"previous_hash"`
		Timestamp    int64          `json:"timestamp"`
		Transactions []*Transaction `json:"transactions"`
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
	return sha256.Sum256([]byte(m))
}

// NewBlock returns a pointer to a new block containing the current timestamp
func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	b := new(Block)
	b.timeStamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	b.transactions = transactions
	return b
}

// LastBlock gets the last block on a blockchain
func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

// CreateBlock creates a new block with the nonce and previousHash and appends it to the current chain
func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	//empties pool
	bc.transactionPool = []*Transaction{}
	bc.chain = append(bc.chain, b)
	return b
}

// Print is a method to improve our block readability
func (b Block) Print() {
	fmt.Printf("timestamp %d\n", b.timeStamp)
	fmt.Printf("previous_hash %x\n", b.previousHash)
	fmt.Printf("nonce %d\n", b.nonce)
	for _, t := range b.transactions {
		t.Print()
	}
}
