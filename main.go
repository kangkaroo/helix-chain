// main.go
package main

import (
	"fmt"
	"helix-chain/consensus/pow"
	"helix-chain/models"
	"helix-chain/storage"
)

func main() {
	ldb, err := storage.NewLevelDB("leveldb")
	if err != nil {
		fmt.Println("Failed to initialize LevelDB:", err)
		return
	}
	defer func(ldb *storage.LevelDB) {
		err := ldb.Close()
		if err != nil {

		}
	}(ldb)
	blockchain, err := models.LoadBlockChainFromDB(ldb)
	if err != nil {
		fmt.Println("Failed to load BlockChain:", err)
	}
	// 如果区块链为空（即没有创世区块），则创建新的区块链
	if nil == blockchain || len(blockchain.Tip) == 0 {
		blockchain, err = models.CreateBlockChainWithGenesisBlock(ldb)
		if err != nil {
			fmt.Println("Failed to create blockchain:", err)
			return
		}
	}

	lastBlock, err := blockchain.GetLastBlock()
	// 添加新块
	data := []byte("Some transaction data")
	newBlock := models.NewBlock(int(lastBlock.Height+1), lastBlock.Hash, data)
	proofOfWork := pow.NewProofOfWork(newBlock)
	hash, nonce := proofOfWork.Run()
	newBlock.Hash = hash
	newBlock.Nonce = nonce
	err = blockchain.AddBlock(newBlock)
	if err != nil {
		fmt.Println("Failed to add new block:", err)
	}

	err = blockchain.PrintAllBlocks()
	if err != nil {
		return
	}
}
