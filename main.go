package main

import (
	"github.com/pedro-git-projects/blockchain/pkg/blockchain"
)

func main() {

	noobChain := blockchain.NewBlockchain()
	noobChain.Print()

	noobChain.AddTransaction("A", "B", 0.32)
	nonce := noobChain.ProofOfWork()
	previousHash := noobChain.LastBlock().Hash()
	noobChain.CreateBlock(nonce, previousHash)

	noobChain.AddTransaction("X", "Y", 4)
	noobChain.AddTransaction("P", "Q", 2.22)
	noobChain.LastBlock().Hash()
	nonce = noobChain.ProofOfWork()
	noobChain.CreateBlock(nonce, previousHash)

	noobChain.Print()

}
