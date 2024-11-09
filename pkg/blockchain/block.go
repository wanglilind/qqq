package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/wanglilind/qqq/pkg/crypto"
)

type Block struct {
	Header BlockHeader   `json:"header"`
	Body   BlockBody     `json:"body"`
	Hash   string        `json:"hash"`
}

type BlockHeader struct {
	Version       uint32    `json:"version"`
	PreviousHash  string    `json:"previous_hash"`
	MerkleRoot    string    `json:"merkle_root"`
	Timestamp     time.Time `json:"timestamp"`
	Height        uint64    `json:"height"`
	ProposerID    string    `json:"proposer_id"`
	Signature     []byte    `json:"signature"`
}

type BlockBody struct {
	Transactions []Transaction `json:"transactions"`
}

func NewBlock(previousHash string, height uint64, proposerID string) *Block {
	block := &Block{
		Header: BlockHeader{
			Version:      1,
			PreviousHash: previousHash,
			Timestamp:    time.Now(),
			Height:      height,
			ProposerID:  proposerID,
		},
		Body: BlockBody{
			Transactions: make([]Transaction, 0),
		},
	}
	
	return block
}

func (b *Block) AddTransaction(tx Transaction) {
	b.Body.Transactions = append(b.Body.Transactions, tx)
	b.Header.MerkleRoot = b.calculateMerkleRoot()
}

func (b *Block) Sign(signer *crypto.Signer) error {
	blockBytes := b.getSigningBytes()
	signature, err := signer.Sign(blockBytes)
	if err != nil {
		return err
	}
	b.Header.Signature = signature
	return nil
}

func (b *Block) CalculateHash() string {
	hasher := sha256.New()
	blockBytes := b.getSigningBytes()
	hasher.Write(blockBytes)
	return hex.EncodeToString(hasher.Sum(nil))
}

func (b *Block) getSigningBytes() []byte {
	// åºåååºåæ°æ®ç¨äºç­¾å?
	// å®éå®ç°éè¦ç¡®ä¿ç¡®å®æ§åºåå
	return []byte{}
}

func (b *Block) calculateMerkleRoot() string {
	// è®¡ç®äº¤æçMerkleæ æ ¹
	// å®éå®ç°éè¦æå»ºå®æ´çMerkleæ ?
	return ""
} 
