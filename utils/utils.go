package utils

import "fmt"

// Log 日志函数
func Log(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
