package node

import (
	consensus "helix-chain/consensus/hotstuff"
	"helix-chain/utils"
	"sync"
	"time"
)

// Node 代表一个节点
type Node struct {
	ID          int
	NumNodes    int
	CurrentView int
	Consensus   *consensus.HotStuff
	mutex       sync.Mutex
}

// NewNode 创建新节点
func NewNode(id int, numNodes int) *Node {
	return &Node{
		ID:          id,
		NumNodes:    numNodes,
		CurrentView: 0,
		Consensus:   consensus.NewHotStuff(id, numNodes),
	}
}

// Start 启动节点
func (n *Node) Start() {
	for {
		// 每个节点定期执行共识
		n.Consensus.ExecuteConsensus()
		time.Sleep(1 * time.Second) // 模拟执行时间
	}
}

// Propose 提交提案
func (n *Node) Propose(command string) {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	n.Consensus.Propose(command)
	utils.Log("Node %d proposed command: %s", n.ID, command)
}
