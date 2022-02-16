package blockchain

import (
	"fmt"
	"strings"

	"github.com/pedro-git-projects/blockchain/pkg/block"
)

var (
	separator = strings.Repeat("-", 35)
)

type Blockchain struct {
	transactionPool []string
	chain           []*block.Block
}

// NewBlockchain returns a pointer to a Blockchain struct containing a new block with a ne Hash
func NewBlockchain() *Blockchain {
	b := &block.Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

// Print is readable print of the blockchain state
func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", separator, i, separator)
		block.Print()
	}
}

// LastBlock gets the last block on a blockchain
func (bc *Blockchain) LastBlock() *block.Block {
	return bc.chain[len(bc.chain)-1]
}

// CreateBlock creates a new block with the nonce and previousHash and appends it to the current chain
func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *block.Block {
	b := block.NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}
