package block

import (
	"fmt"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timeStamp    int64
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.timeStamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}

func (b Block) Print() {
	fmt.Printf("timestamp %d\n", b.timeStamp)
	fmt.Printf("previous_hash %s\n", b.previousHash)
	fmt.Printf("nonce %d\n", b.nonce)
	fmt.Printf("transactions %v\n", b.transactions)
}
