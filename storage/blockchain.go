package storage

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
}

type Blockchain struct {
	Blocks []Block
}

// 初始化区块链
func InitBlockchain() *Blockchain {
	return &Blockchain{Blocks: []Block{{Index: 0}}}
}

// 添加区块
func (bc *Blockchain) AddBlock(data string) {
	// 实现区块添加逻辑
}
