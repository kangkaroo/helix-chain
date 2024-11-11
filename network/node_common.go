package network

import (
	"sync"
)

// AllNodeCommonMsg 包含所有节点的通用信息
type AllNodeCommonMsg struct {
	mu                sync.RWMutex
	AllNodeAddressMap map[int]BasicPeer // 保存节点的 IP 地址和端口号
	View              int               // 当前视图的编号，0 表示未初始化
	Size              int               // 区块链中的节点总数
}

// NewAllNodeCommonMsg 初始化 AllNodeCommonMsg
func NewAllNodeCommonMsg() *AllNodeCommonMsg {
	return &AllNodeCommonMsg{
		AllNodeAddressMap: make(map[int]BasicPeer),
		View:              0,
		Size:              1, // 默认值，表示只有一个节点
	}
}

// GetMaxF 计算最大失效节点数量
func (m *AllNodeCommonMsg) GetMaxF() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return (m.Size - 1) / 3
}

// GetPriIndex 获取主节点的序号
func (m *AllNodeCommonMsg) GetPriIndex() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return (m.View + 1) % m.Size
}

// AddNode 添加新节点到 AllNodeAddressMap，并更新 Size
func (m *AllNodeCommonMsg) AddNode(index int, info BasicPeer) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.AllNodeAddressMap[index] = info
	m.Size = len(m.AllNodeAddressMap) + 1
}

// UpdateView 更新视图编号
func (m *AllNodeCommonMsg) UpdateView(newView int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.View = newView
}
