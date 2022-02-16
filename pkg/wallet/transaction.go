package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"
)

type Transaction struct {
	senderPrivateKey          *ecdsa.PrivateKey
	senderPublicKey           *ecdsa.PublicKey
	senderBlockChainAdress    string
	recipientBlockChainAdress string
	value                     float32
}

type Signature struct {
	R *big.Int
	S *big.Int
}

func NewTransaction(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, sender string, recipient string, value float32) *Transaction {
	return &Transaction{privateKey, publicKey, sender, recipient, value}
}

func (t *Transaction) GenerateSignature() *Signature {
	m, _ := json.Marshal(t)
	h := sha256.Sum256([]byte(m))

	r, s, _ := ecdsa.Sign(rand.Reader, t.senderPrivateKey, h[:])
	return &Signature{r, s}
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBlockChainAdress,
		Recipient: t.recipientBlockChainAdress,
		Value:     t.value,
	})
}

func (s *Signature) String() string {
	return fmt.Sprintf("%x%x", s.R, s.S)
}
