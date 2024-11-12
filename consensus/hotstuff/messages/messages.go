package messages

// Message 消息结构体
type Message struct {
	Type    string // 消息类型
	NodeID  int    // 节点 ID
	View    int    // 视图号
	Command string // 提交的指令
}
