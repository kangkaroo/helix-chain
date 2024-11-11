package pow

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"helix-chain/storage"
	"math"
	"math/big"
)

const targetBits = 24 // 目标难度的位数

type ProofOfWork struct {
	Block  *storage.Block // 需要共识验证的区块
	target *big.Int       // 目标难度的哈希
}

// NewProofOfWork 构建 ProofOfWork 实例
func NewProofOfWork(block *storage.Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits)) // 设置目标难度
	pow := &ProofOfWork{block, target}
	return pow
}

// Run 执行工作量证明的挖矿过程
func (pow *ProofOfWork) Run() ([]byte, int64) {
	var hashInt big.Int
	var hash [32]byte
	nonce := int64(0)

	for nonce < math.MaxInt64 {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		// 检查哈希值是否小于目标难度
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}

	return hash[:], nonce
}

// prepareData 准备要进行哈希计算的数据
func (pow *ProofOfWork) prepareData(nonce int64) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PreBlockHash,
			pow.Block.Data,
			IntToHex(pow.Block.TimeStamp),
			IntToHex(pow.Block.Height),
			IntToHex(nonce),
		},
		[]byte{},
	)

	return data
}

// IntToHex 将整数转换为十六进制表示
func IntToHex(n int64) []byte {
	buff := new(bytes.Buffer)
	_ = binary.Write(buff, binary.BigEndian, n)
	return buff.Bytes()
}
