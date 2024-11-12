// pow/pow.go
package pow

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"helix-chain/models"
	"math"
	"math/big"
)

const targetBits = 24

type ProofOfWork struct {
	Block  *models.Block
	target *big.Int
}

func NewProofOfWork(block *models.Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	return &ProofOfWork{block, target}
}

func (pow *ProofOfWork) Run() ([]byte, int64) {
	var hashInt big.Int
	var hash [32]byte
	nonce := int64(0)

	for nonce < math.MaxInt64 {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}

	return hash[:], nonce
}

func (pow *ProofOfWork) prepareData(nonce int64) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PreBlockHash,
			pow.Block.Data,
			IntToHex(pow.Block.Timestamp),
			IntToHex(pow.Block.Height),
			IntToHex(nonce),
		},
		[]byte{},
	)

	return data
}

func IntToHex(n int64) []byte {
	buff := new(bytes.Buffer)
	_ = binary.Write(buff, binary.BigEndian, n)
	return buff.Bytes()
}
