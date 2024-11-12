package models

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"time"
)

// Block structure
type Block struct {
	Height       int64
	PreBlockHash []byte
	Data         []byte
	Hash         []byte
	Nonce        int64
	Timestamp    int64 // Add a Timestamp field
}

// NewBlock - creates a new block
func NewBlock(height int, preBlockHash, data []byte) *Block {
	block := &Block{
		Height:       int64(height),
		PreBlockHash: preBlockHash,
		Data:         data,
		Timestamp:    time.Now().Unix(), // Set the current Unix timestamp
	}
	return block
}

// calculateHash - generates hash for the block, including height and timestamp
func (b *Block) calculateHash() []byte {
	data := bytes.Join([][]byte{
		[]byte(fmt.Sprintf("%d", b.Height)),
		b.PreBlockHash,
		b.Data,
		[]byte(fmt.Sprintf("%d", b.Nonce)),
		[]byte(fmt.Sprintf("%d", b.Timestamp)), // Include Timestamp in hash calculation
	}, []byte{})

	hash := sha256.Sum256(data)
	return hash[:]
}
