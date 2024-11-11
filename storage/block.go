package storage

import (
	"bytes"
	"fmt"
	"helix-chain/common/utils"
	"time"
)

type Block struct {
	TimeStamp    int64  // 当前时间戳
	Hash         []byte // 当前区块的哈希值
	PreBlockHash []byte // 上一个区块的哈希值
	Height       int64  // 区块高度
	Data         []byte // 区块数据
}

// NewBlock 创建一个新的区块
func NewBlock(height int64, preBlockHash []byte, data []byte) *Block {
	block := Block{
		TimeStamp:    time.Now().UnixMilli(),
		Hash:         nil,
		PreBlockHash: preBlockHash,
		Height:       height,
		Data:         data,
	}
	block.Hash = block.calculateHash()
	return &block
}

// calculateHash 计算区块的哈希值
func (b *Block) calculateHash() []byte {
	data := bytes.Join([][]byte{
		b.PreBlockHash,
		b.Data,
		[]byte(fmt.Sprintf("%d", b.TimeStamp)),
		[]byte(fmt.Sprintf("%d", b.Height)),
	}, []byte{})
	return utils.Hash256(data)
}

// ToString 格式化区块信息为字符串
func (b *Block) ToString() string {
	return fmt.Sprintf("Block{Height: %d, Data: %s, TimeStamp: %d, Hash: %x, PreBlockHash: %x}",
		b.Height, string(b.Data), b.TimeStamp, b.Hash, b.PreBlockHash)
}

// CreateGenesisBlock 创建创世区块
func CreateGenesisBlock(data []byte) *Block {
	return NewBlock(1, nil, data)
}

// CreateBlockChainWithGenesisBlock 创建一个带创世区块的区块链
func CreateBlockChainWithGenesisBlock(data []byte) *BlockChain {
	genesisBlock := CreateGenesisBlock(data)
	return &BlockChain{Blocks: []*Block{genesisBlock}}
}
