package main

import (
	"github.com/pedro-git-projects/blockchain/pkg/blockchain"
)

func main() {
	noobChain := blockchain.NewBlockchain()
	noobChain.CreateBlock(9, "hash 1")
	noobChain.CreateBlock(12, "hash 2")
	noobChain.Print()
}
