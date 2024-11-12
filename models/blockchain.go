// models/blockchain.go
package models

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"helix-chain/storage"
)

// BlockChain 结构体
type BlockChain struct {
	ldb *storage.LevelDB
	Tip []byte
}

// CreateBlockChainWithGenesisBlock 初始化区块链并创建创世区块
func CreateBlockChainWithGenesisBlock(ldb *storage.LevelDB) (*BlockChain, error) {
	exists, err := ldb.Has([]byte("l"))
	if err != nil {
		return nil, err
	}

	var tip []byte
	if !exists {
		genesisBlock := &Block{Height: 0, PreBlockHash: nil, Data: []byte("Genesis Block")}
		genesisBlock.Hash = genesisBlock.calculateHash()

		err := ldb.Put(genesisBlock.Hash, serializeBlock(genesisBlock))
		if err != nil {
			return nil, err
		}

		err = ldb.Put([]byte("l"), genesisBlock.Hash)
		if err != nil {
			return nil, err
		}

		tip = genesisBlock.Hash
	} else {
		tip, err = ldb.Get([]byte("l"))
		if err != nil {
			return nil, err
		}
	}

	return &BlockChain{ldb: ldb, Tip: tip}, nil
}

// 从数据库加载区块链
func LoadBlockChainFromDB(ldb *storage.LevelDB) (*BlockChain, error) {
	tip, err := ldb.Get([]byte("l"))
	if err != nil {
		return nil, fmt.Errorf("no blockchain found, creating new one")
	}

	return &BlockChain{ldb: ldb, Tip: tip}, nil
}

// AddBlock - adds a new block to the blockchain
func (bc *BlockChain) AddBlock(newBlock *Block) error {
	// 将新区块序列化并存储到数据库
	err := bc.ldb.Put(newBlock.Hash, serializeBlock(newBlock))
	if err != nil {
		return err
	}
	// 更新 Tip 为新区块的哈希
	err = bc.ldb.Put([]byte("l"), newBlock.Hash)
	if err != nil {
		return err
	}
	// 更新区块链的 Tip
	bc.Tip = newBlock.Hash
	return nil
}

// 序列化和反序列化方法
func serializeBlock(block *Block) []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(block)
	if err != nil {
		fmt.Println("Error serializing block:", err)
		return nil
	}
	return result.Bytes()
}

func deserializeBlock(data []byte) (*Block, error) {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil {
		return nil, err
	}
	return &block, nil
}

// 获取区块链中的最后一个区块
func (bc *BlockChain) GetLastBlock() (*Block, error) {
	// 从数据库获取最后一个区块的哈希值
	lastBlockData, err := bc.ldb.Get(bc.Tip)
	if err != nil {
		return nil, fmt.Errorf("failed to get the last block: %v", err)
	}
	// 反序列化获取到的区块数据
	lastBlock, err := deserializeBlock(lastBlockData)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize the last block: %v", err)
	}

	return lastBlock, nil
}

// 查询区块
func (bc *BlockChain) GetBlock(hash []byte) (*Block, error) {
	blockData, err := bc.ldb.Get(hash)
	if err != nil {
		return nil, err
	}

	return deserializeBlock(blockData)
}
