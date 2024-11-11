package main

import (
	"fmt"
	"helix-chain/consensus/pow"
	"helix-chain/storage"
)

func main() {
	// 创建带有创世区块的区块链
	blockChain := storage.CreateBlockChainWithGenesisBlock([]byte("GenesisBlock"))

	// 定义要添加的区块数据
	blockData := []string{"Block 1 Data", "Block 2 Data", "Block 3 Data"}

	// 添加多个区块到区块链
	for _, data := range blockData {
		// 获取当前链的最后一个区块
		prevBlock := blockChain.Blocks[len(blockChain.Blocks)-1]

		// 创建新的区块
		newBlock := storage.NewBlock(prevBlock.Height+1, prevBlock.Hash, []byte(data))

		// 使用 ProofOfWork 生成区块的哈希
		pow := pow.NewProofOfWork(newBlock)
		hash, nonce := pow.Run()

		// 将工作量证明结果应用到区块
		newBlock.Hash = hash
		fmt.Printf("Mined Block %d with nonce %d: %x\n", newBlock.Height, nonce, newBlock.Hash)

		// 将新块添加到区块链
		blockChain.AddBlock([]byte(data))
	}

	// 打印区块链中的所有区块信息
	for _, block := range blockChain.Blocks {
		fmt.Println(block.ToString())
	}
}
