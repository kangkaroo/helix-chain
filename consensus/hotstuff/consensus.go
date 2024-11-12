package consensus

import (
	"helix-chain/consensus/hotstuff/messages"
	"helix-chain/utils" // 确保引入的工具包正确
	"time"
)

// HotStuff 结构体表示 HotStuff 共识协议
type HotStuff struct {
	ID                  int
	CurrentView         int
	PrepareMsg          *PrepareMessage // 更具描述性的消息结构体
	CollectedSignatures map[int]bool    // 用于存储签名的集合
	RequiredSigs        int             // 收集所需的签名数
	SignatureChannel    chan int        // 更具描述性的通道名称
}

// PrepareMessage 表示准备阶段消息结构
type PrepareMessage struct {
	Command string
	NodeID  int
	View    int
}

// NewHotStuffConsensus 初始化 HotStuff 协议
func NewHotStuffConsensus(id, requiredSigs int) *HotStuff {
	return &HotStuff{
		ID:                  id,
		CurrentView:         0,
		CollectedSignatures: make(map[int]bool),
		RequiredSigs:        requiredSigs,
		SignatureChannel:    make(chan int), // 创建签名通道
	}
}

// ReceivePrepareMessage 处理准备阶段消息
func (h *HotStuff) ReceivePrepareMessage(msg *PrepareMessage) {
	// 验证消息的有效性
	if msg.NodeID <= 0 || msg.View != h.CurrentView {
		utils.Log("Node %d received invalid prepare message from Node %d", h.ID, msg.NodeID)
		return
	}

	// 当收到准备消息时，向 SignatureChannel 发送节点 ID
	select {
	case h.SignatureChannel <- msg.NodeID: // 发送节点 ID
		utils.Log("Node %d collected prepare message from Node %d", h.ID, msg.NodeID)
	default:
		utils.Log("Node %d could not send prepare message from Node %d to channel", h.ID, msg.NodeID)
	}
}

// PreparePhase 准备阶段
func (h *HotStuff) PreparePhase() {
	utils.Log("Node %d in Prepare Phase", h.ID)

	// 广播准备消息
	h.Broadcast(h.PrepareMsg)

	// 设置超时处理逻辑
	timeout := time.After(5 * time.Second) // 设置 5 秒超时

	// 启动一个 goroutine 监听其他节点的准备消息
	go func() {
		for {
			select {
			case sigNodeID := <-h.SignatureChannel:
				h.CollectedSignatures[sigNodeID] = true // 收集签名

				// 检查是否收集到足够的签名
				if len(h.CollectedSignatures) >= h.RequiredSigs {
					h.PreCommitPhase() // 进入预提交阶段
					return
				}
			case <-timeout:
				utils.Log("Node %d timed out waiting for prepare messages", h.ID)
				return
			}
		}
	}()
}

// PreCommitPhase 预提交阶段
func (h *HotStuff) PreCommitPhase() {
	utils.Log("Node %d in Pre-Commit Phase", h.ID)

	// 生成 PreCommit 消息
	preCommitMsg := &messages.Message{
		Type:    "PreCommit",
		NodeID:  h.ID,
		View:    h.CurrentView,
		Command: h.ProposeMsg.Command,
	}

	// 广播预提交消息
	h.Broadcast(preCommitMsg)

	// 这里可以添加超时处理逻辑
	// 等待接收来自其他节点的预提交消息并收集签名
	// 假设收集了足够的签名后，进入提交阶段
	h.CommitPhase()
}

// CommitPhase 提交阶段
func (h *HotStuff) CommitPhase() {
	utils.Log("Node %d in Commit Phase", h.ID)

	// 生成 Commit 消息
	commitMsg := &messages.Message{
		Type:    "Commit",
		NodeID:  h.ID,
		View:    h.CurrentView,
		Command: h.ProposeMsg.Command,
	}

	// 广播提交消息
	h.Broadcast(commitMsg)

	// 执行成功提交逻辑
	h.ExecuteCommand(h.ProposeMsg.Command)

	// 更新视图，如果需要的话
	h.CurrentView++ // 这里假设视图增加的条件已经满足
}

// Broadcast 广播消息
func (h *HotStuff) Broadcast(msg *messages.Message) {
	// 广播消息逻辑，发送给其他节点
	utils.Log("Node %d broadcasting message: %s", h.ID, msg.Type)
}

// ListenForMessages 监听消息
func (h *HotStuff) ListenForMessages() {
	for msg := range h.Messages {
		h.ProcessMessage(msg)
	}
}

// ProcessMessage 处理接收到的消息
func (h *HotStuff) ProcessMessage(msg *messages.Message) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	utils.Log("Node %d received message: %s from Node %d", h.ID, msg.Type, msg.NodeID)

	// 根据消息类型处理
	switch msg.Type {
	case "Prepare":
		// 收集准备消息，验证并更新准备证书
		h.PrepareQC = &QC{View: msg.View, NodeID: msg.NodeID}
		// 检查是否可以进入预提交阶段
	case "PreCommit":
		// 收集预提交消息，验证并更新预提交证书
		h.PreCommitQC = &QC{View: msg.View, NodeID: msg.NodeID}
		// 检查是否可以进入提交阶段
	case "Commit":
		// 收集提交消息，最终确认
		h.CommitQC = &QC{View: msg.View, NodeID: msg.NodeID}
		// 可以进行后续处理
	}
}

// ExecuteCommand 执行指令
func (h *HotStuff) ExecuteCommand(command string) {
	// 这里可以添加执行命令的具体逻辑
	utils.Log("Node %d executing command: %s", h.ID, command)
}
