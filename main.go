package main

import (
	"fmt"

	"github.com/pedro-git-projects/blockchain/pkg/wallet"
)

func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())
	fmt.Println(w.BlockchainAddress())

	t := wallet.NewTransaction(w.PrivateKey(), w.PublicKey(), w.BlockchainAddress(), "B", 1.00)
	fmt.Println("Signature", t.GenerateSignature())
}
