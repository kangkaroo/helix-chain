package consensus

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

// PbftMsg 结构体
type PbftMsg struct {
	MsgType int    `json:"msgType"` // 消息类型
	Body    string `json:"body"`    // 消息体
	Node    int    `json:"node"`    // 消息发起的结点编号
	ToNode  int    `json:"toNode"`  // 消息发送的目的地
	Time    int64  `json:"time"`    // 消息时间戳
	IsOk    bool   `json:"isOk"`    // 检测是否通过
	ViewNum int    `json:"viewNum"` // 结点视图
	ID      string `json:"id"`      // 使用UUID进行生成
}

// 构造函数
func NewPbftMsg(msgType, node int) *PbftMsg {
	return &PbftMsg{
		MsgType: msgType,
		Node:    node,
		Time:    time.Now().UnixMilli(), // 获取当前时间戳（毫秒）
		ID:      uuid.New().String(),    // 生成UUID
		ViewNum: AllNodeCommonMsgView,   // 假设 AllNodeCommonMsg.view 赋值
	}
}

// 重写 equals 方法
func (msg *PbftMsg) Equals(o *PbftMsg) bool {
	if o == nil {
		return false
	}
	return msg.Node == o.Node &&
		msg.Time == o.Time &&
		msg.ViewNum == o.ViewNum &&
		msg.Body == o.Body &&
		msg.ID == o.ID
}

// 重写 hashCode 方法
func (msg *PbftMsg) HashCode() int {
	// 这里只是一个简单示例，可以根据需求调整hash算法
	return int(msg.Time) ^ msg.Node ^ msg.ViewNum
}

// 假设 AllNodeCommonMsgView 是一个全局常量
var AllNodeCommonMsgView = 1

func main() {
	// 创建一个新的 PbftMsg 实例
	msg := NewPbftMsg(1, 100)
	fmt.Println(msg)

	// 示例：比较两个 PbftMsg 实例
	msg2 := NewPbftMsg(1, 100)
	fmt.Println(msg.Equals(msg2)) // 输出：true
}
