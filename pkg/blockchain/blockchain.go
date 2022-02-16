package blockchain

import (
	"fmt"
	"strings"

	"github.com/pedro-git-projects/blockchain/pkg/block"
)

var (
	separator = strings.Repeat("#", 25)
)

type Blockchain struct {
	transactionPool []string
	chain           []*block.Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Initial hash")
	return bc
}

func (bc Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", separator, i, separator)
		block.Print()
	}
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *block.Block {
	b := block.NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}
