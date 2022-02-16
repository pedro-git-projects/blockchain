// Package transaction contains types functions and methods associated with blockchain transactions
package transaction

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
