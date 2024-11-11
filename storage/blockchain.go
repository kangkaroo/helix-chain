package storage

// BlockChain 表示一个简单的区块链结构
type BlockChain struct {
	Blocks []*Block // 用于存储区块链中区块的切片
}

// AddBlock 向区块链添加一个包含指定数据的新区块
func (blockChain *BlockChain) AddBlock(data []byte) {
	// 获取区块链中的最后一个区块
	prevBlock := blockChain.Blocks[len(blockChain.Blocks)-1]
	// 使用上一个区块的哈希和高度 + 1 创建一个新区块
	newBlock := NewBlock(prevBlock.Height+1, prevBlock.Hash, data)
	// 将新区块添加到区块链中
	blockChain.Blocks = append(blockChain.Blocks, newBlock)
}
