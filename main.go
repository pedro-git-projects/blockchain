package main

import (
	"github.com/pedro-git-projects/blockchain/pkg/blockchain"
)

func main() {
	// first block
	noobChain := blockchain.NewBlockchain()

	previousHash := noobChain.LastBlock().Hash()
	// creating second block storing previous Hash
	noobChain.CreateBlock(7, previousHash)

	previousHash = noobChain.LastBlock().Hash()
	// creating second block storing previous Hash
	noobChain.CreateBlock(3, previousHash)

	noobChain.Print()

}
