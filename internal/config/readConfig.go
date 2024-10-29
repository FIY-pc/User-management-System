package config

import (
	"encoding/json"
	"log"
	"os"
)

// Config 所有配置信息将从json文件加载到该结构体实例中
var Config configStruct

// configStruct 配置信息结构体
type configStruct struct {
	Server   ServerConfig   `json:"server"`
	Jwt      JwtConfig      `json:"jwt"`
	Admin    AdminConfig    `json:"admin"`
	Postgres PostgresConfig `json:"postgresql"`
	Bcrypt   BcryptConfig   `json:"bcrypt"`
}

// ServerConfig 服务配置信息
type ServerConfig struct {
	Port      string `json:"port"`
	ApiPrefix string `json:"api_prefix"`
	TestURL   string `json:"test_url"`
}

// JwtConfig jwt配置信息
type JwtConfig struct {
	Secret string `json:"secret"`
	Exp    int64  `json:"exp"`
}

// AdminConfig 管理员配置信息
type AdminConfig struct {
	AdminName string `json:"adminName"`
	AdminPass string `json:"adminPass"`
}

// PostgresConfig postgresql配置信息
type PostgresConfig struct {
	Dsn string `json:"dsn"`
}

// BcryptConfig bcrypt配置信息
type BcryptConfig struct {
	Cost int `json:"cost"`
}

// InitConfig 初始化配置信息
func InitConfig() {
	data, err := os.ReadFile("./Config/default.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &Config)
	if err != nil {
		panic(err)
	}
	log.Println("read local Config success")
}
