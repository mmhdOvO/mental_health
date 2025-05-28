// config/config.go
package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Config 定义全局配置结构
type Config struct {
	Database DatabaseConfig `yaml:"database"` // 数据库配置
	Server   ServerConfig   `yaml:"server"`   // 服务器配置
	JWT      JWTConfig      `yaml:"jwt"`      // JWT 配置（可选）
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	DSN string `yaml:"dsn"` // 数据库连接字符串
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port string `yaml:"port"` // 监听端口
}

// JWTConfig JWT 配置（可选）
type JWTConfig struct {
	Secret      string `yaml:"secret"`       // 签名密钥
	ExpireHours int    `yaml:"expire_hours"` // 有效期（小时）
}

// LoadConfig 加载并解析配置文件（返回错误而非直接终止）
func LoadConfig(path string) (*Config, error) {
	// 读取配置文件内容
	data, err := ioutil.ReadFile(path)
	if err != nil {
		// 返回包装后的错误，不直接调用 log.Fatalf
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 解析 YAML 到结构体
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	return &cfg, nil
}
