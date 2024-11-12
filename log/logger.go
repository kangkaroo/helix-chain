package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

// Logger 是全局日志实例
var Logger *logrus.Logger

// 初始化日志配置
func InitLogger() {
	Logger = logrus.New()

	// 设置日志输出为标准输出
	Logger.SetOutput(os.Stdout)

	// 设置日志级别，可以根据需要选择 debug, info, warn, error, fatal, panic
	Logger.SetLevel(logrus.DebugLevel)

	// 设置日志格式为文本格式
	Logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,                  // 强制使用颜色
		FullTimestamp:   true,                  // 显示完整的时间戳
		TimestampFormat: "2006-01-02 15:04:05", // 时间格式
	})
}

// 记录调试日志
func Debug(args ...interface{}) {
	Logger.Debug(args...)
}

// 记录信息日志
func Info(args ...interface{}) {
	Logger.Info(args...)
}

// 记录警告日志
func Warn(args ...interface{}) {
	Logger.Warn(args...)
}

// 记录错误日志
func Error(args ...interface{}) {
	Logger.Error(args...)
}

// 记录致命错误日志
func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}

// 记录恐慌日志
func Panic(args ...interface{}) {
	Logger.Panic(args...)
}
