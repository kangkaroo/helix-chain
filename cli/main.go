package main

import (
	"flag"
	"fmt"
	"helix-chain/consensus/pow"
	"helix-chain/models"
	"helix-chain/storage"
	"os"
)

func main() {
	// 设置命令行参数
	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	queryBlockCmd := flag.NewFlagSet("query", flag.ExitOnError)

	// 添加区块命令参数
	addData := addBlockCmd.String("data", "", "Data for the new block")

	// 查询区块命令参数
	queryHash := queryBlockCmd.String("hash", "", "Hash of the block to query")

	// 解析命令行参数
	if len(os.Args) < 2 {
		fmt.Println("expected 'add' or 'query' subcommands")
		os.Exit(1)
	}

	// 创建一个 LevelDB 实例
	ldb, err := storage.NewLevelDB("blockchain.db")
	if err != nil {
		fmt.Println("Error opening LevelDB:", err)
		os.Exit(1)
	}
	defer ldb.Close()

	// 加载或创建区块链
	blockchain, err := models.LoadBlockChainFromDB(ldb)
	if err != nil {
		fmt.Println("Failed to load BlockChain:", err)
	}
	// 如果区块链为空（即没有创世区块），则创建新的区块链
	if blockchain == nil || len(blockchain.Tip) == 0 {
		blockchain, err = models.CreateBlockChainWithGenesisBlock(ldb)
		if err != nil {
			fmt.Println("Failed to create blockchain:", err)
			return
		}
	}

	lastBlock, err := blockchain.GetLastBlock()
	if err != nil {
		fmt.Println("Error retrieving last block:", err)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		// 解析添加区块命令
		addBlockCmd.Parse(os.Args[2:])
		if *addData == "" {
			fmt.Println("You must provide data for the block using -data flag")
			os.Exit(1)
		}

		// 创建新区块
		newBlock := models.NewBlock(int(lastBlock.Height+1), lastBlock.Hash, []byte(*addData))
		proofOfWork := pow.NewProofOfWork(newBlock)
		hash, nonce := proofOfWork.Run()
		newBlock.Hash = hash
		newBlock.Nonce = nonce

		// 添加新区块到区块链
		err = blockchain.AddBlock(newBlock)
		if err != nil {
			fmt.Println("Error adding block:", err)
		} else {
			fmt.Println("Block added successfully")
		}

	case "query":
		// 解析查询区块命令
		queryBlockCmd.Parse(os.Args[2:])
		if *queryHash == "" {
			// 查询最新区块
			block, err := blockchain.GetLastBlock()
			if err != nil {
				fmt.Println("Error retrieving last block:", err)
			} else {
				// 打印区块信息
				fmt.Printf("Last Block: Height: %d, Hash: %x, Nonce: %d, PrevHash: %x, Data: %s\n",
					block.Height, block.Hash, block.Nonce, block.PreBlockHash, block.Data)
			}
		} else {
			// 根据哈希查询特定区块
			block, err := blockchain.GetBlock([]byte(*queryHash))
			if err != nil {
				fmt.Println("Error retrieving block:", err)
			} else {
				// 打印区块信息
				fmt.Printf("Block found: Height: %d, Hash: %x, Nonce: %d, PrevHash: %x, Data: %s\n",
					block.Height, block.Hash, block.Nonce, block.PreBlockHash, block.Data)
			}
		}

	default:
		fmt.Println("expected 'add' or 'query' subcommands")
		os.Exit(1)
	}
}
