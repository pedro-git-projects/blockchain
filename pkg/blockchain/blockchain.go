// Package blockchain contains types, functions and methods for a proof of work blockchain structure
package blockchain

import (
	"fmt"
	"strings"
)

var (
	separatorBlock = strings.Repeat("=", 36)
	separatorEnd   = strings.Repeat("*", 81)
)

// Blockchain contains a pointer to a transaction pool, pointer to a block and an address
type Blockchain struct {
	transactionPool   []*Transaction
	chain             []*Block
	blockchainAddress string
}

// NewBlockchain takes a address value and returns a pointer to a Blockchain struct containing a new block with a new hash
func NewBlockchain(blockchainAddress string) *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.blockchainAddress = blockchainAddress
	bc.CreateBlock(0, b.Hash())
	return bc
}

// Print overloads the Print function and makes so that the struct is more readable
func (bc Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", separatorBlock, i, separatorBlock)
		block.Print()
	}
	fmt.Println(separatorEnd)
}
