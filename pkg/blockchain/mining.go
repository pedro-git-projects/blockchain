package blockchain

import (
	"fmt"
	"log"
	"strings"
)

const (
	MINING_DIFFICULTY = 3
	MINING_SENDER     = "THE BLOCKCHAIN"
	MINING_REWARD     = 1.00
)

// ValidProof will detrime how many prefix positions in our hash are zeroes
func (bc *Blockchain) ValidProof(nonce int, previousHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeroes := strings.Repeat("0", difficulty)
	guessBlock := Block{0, nonce, previousHash, transactions}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficulty] == zeroes
}

func (bc *Blockchain) ProofOfWork() int {
	// copying the contents of the transaction pool
	transactions := bc.CopyTransactionPool()
	// refering to the last Hash on the chain
	previousHash := bc.LastBlock().Hash()
	var nonce int
	for !bc.ValidProof(nonce, previousHash, transactions, MINING_DIFFICULTY) {
		// increment till our hash contains all necessary zeroes
		nonce += 1
	}
	return nonce
}

func (bc *Blockchain) Mining() bool {
	// When mining is completed successfully, the mining reward of 1.00 is sent to the specified addres
	bc.AddTransaction(MINING_SENDER, bc.blockchainAddress, MINING_REWARD)
	// Compute the nonce by the POW algorithm with the reward already included
	nonce := bc.ProofOfWork()
	previousHash := bc.LastBlock().Hash()
	// Then we create the new block
	bc.CreateBlock(nonce, previousHash)
	log.Println("action=mining, status=success")
	return true
}
