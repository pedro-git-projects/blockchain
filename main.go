package main

import (
	"github.com/pedro-git-projects/blockchain/pkg/blockchain"
)

func main() {

	myAddress := "my_address"
	noobChain := blockchain.NewBlockchain(myAddress)

	noobChain.AddTransaction("A", "B", 0.32)
	noobChain.Mining()

	noobChain.AddTransaction("X", "Y", 4)
	noobChain.AddTransaction("P", "Q", 2.22)
	noobChain.Mining()

	noobChain.Print()

}
