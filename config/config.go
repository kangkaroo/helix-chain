package config

// Config 配置结构体
type Config struct {
	NumNodes int
}

// LoadConfig 加载配置
func LoadConfig() *Config {
	return &Config{
		NumNodes: 4, // 设定节点数量
	}
}
