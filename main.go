package main

import (
	"helix-chain/api"
	"helix-chain/network"
	"log"
	"net/http"
)

func main() {
	// 初始化网络模块
	network.InitNode()

	// 启动 API 服务
	http.HandleFunc("/api", api.HandleAPI)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
