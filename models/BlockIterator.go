package models

import (
	"fmt"
	"helix-chain/storage"
)

// BlockIterator 用于遍历区块链
type BlockIterator struct {
	ldb     *storage.LevelDB
	current []byte // 当前区块的哈希值
	hasNext bool   // 标记是否还有下一个区块
}

// NewBlockIterator 创建一个新的区块链迭代器
func NewBlockIterator(ldb *storage.LevelDB, tip []byte) *BlockIterator {
	return &BlockIterator{
		ldb:     ldb,
		current: tip,
		hasNext: true, // 初始时认为有下一个区块
	}
}

// Next 返回下一个区块，并更新当前区块的哈希值
func (it *BlockIterator) Next() (*Block, error) {
	if !it.hasNext {
		return nil, fmt.Errorf("no more blocks")
	}

	// 从数据库中获取当前区块
	blockData, err := it.ldb.Get(it.current)
	if err != nil {
		return nil, err
	}

	// 反序列化区块
	block, err := deserializeBlock(blockData)
	if err != nil {
		return nil, err
	}

	// 更新当前区块为前一个区块
	it.current = block.PreBlockHash

	// 如果没有前一个区块，则说明遍历完了，设置 hasNext 为 false
	if len(it.current) == 0 {
		it.hasNext = false
	}

	return block, nil
}

// HasNext 判断是否还有下一个区块
func (it *BlockIterator) HasNext() bool {
	return it.hasNext
}

func (bc *BlockChain) PrintAllBlocks() error {
	// 创建迭代器
	iterator := NewBlockIterator(bc.ldb, bc.Tip)

	// 遍历区块链中的所有区块
	for iterator.HasNext() {
		block, err := iterator.Next()
		if err != nil {
			return err
		}

		// 打印当前区块的信息
		fmt.Printf("Height: %d, Hash: %x,Nonce: %d, PrevHash: %x, Data: %s\n", block.Height, block.Hash, block.Nonce, block.PreBlockHash, block.Data)
	}

	return nil
}
