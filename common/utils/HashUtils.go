package utils

import (
	"crypto/sha256"
)

/*
计算hash
*/
// hash256 计算输入字符串的 SHA-256 哈希值，并返回十六进制字符串
func Hash256(data []byte) []byte {
	// 计算 SHA-256 哈希值
	hash := sha256.Sum256(data)
	// 将哈希值转换为十六进制字符串并返回
	return hash[:]
}
