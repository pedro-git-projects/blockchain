package blockchain

import (
	"encoding/json"
	"fmt"
	"strings"
)

var (
	separator = strings.Repeat("-", 79)
)

type Transaction struct {
	senderBlockChainAdress    string
	recipientBlockChainAdress string
	value                     float32
}

// NewTransaction creates a transaction struct and populates it with the sender, recipient and value values
func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", separator)
	fmt.Printf("sender_blockchain_address %s\n", t.senderBlockChainAdress)
	fmt.Printf("recipient_blockchain_address %s\n", t.recipientBlockChainAdress)
	fmt.Printf("value %.2f\n", t.value)
}

// Overloading mashaling method
func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipientBlockChainAdress"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBlockChainAdress,
		Recipient: t.recipientBlockChainAdress,
		Value:     t.value,
	})
}

func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}

func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, t := range bc.transactionPool {
		transactions = append(transactions, NewTransaction(t.senderBlockChainAdress, t.recipientBlockChainAdress, t.value))
	}
	return transactions
}

// CalculateTotalAmount calculates transaction results
func (bc *Blockchain) CalculateTotalAmount(blockchainAddress string) float32 {
	var totalAmount float32
	for _, b := range bc.chain {
		for _, t := range b.transactions {
			value := t.value
			if blockchainAddress == t.recipientBlockChainAdress {
				totalAmount += value
			}

			if blockchainAddress == t.senderBlockChainAdress {
				totalAmount -= value
			}
		}
	}
	return totalAmount
}
