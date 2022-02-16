package main

import (
	"fmt"

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

	fmt.Printf("my_address %2.f\n", noobChain.CalculateTotalAmount("my_address"))
	fmt.Printf("X %2.f\n", noobChain.CalculateTotalAmount("X"))
	fmt.Printf("Y %2.f\n", noobChain.CalculateTotalAmount("Y"))
}
