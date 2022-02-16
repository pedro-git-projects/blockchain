// Package blockchain provides the blockchain type and functions
package blockchain

import (
	"fmt"
	"strings"
)

var (
	separatorBlock = strings.Repeat("=", 35)
)

type Blockchain struct {
	transactionPool   []*Transaction
	chain             []*Block
	blockchainAddress string
}

// NewBlockchain returns a pointer to a Blockchain struct containing a new block with a ne Hash
func NewBlockchain(blockchainAddress string) *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.blockchainAddress = blockchainAddress
	bc.CreateBlock(0, b.Hash())
	return bc
}

// Print is readable print of the blockchain state
func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", separatorBlock, i, separatorBlock)
		block.Print()
	}
}
