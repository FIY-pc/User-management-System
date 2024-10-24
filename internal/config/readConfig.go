package config

import (
	"encoding/json"
	"log"
	"os"
)

var Config configStruct

type configStruct struct {
	Server   ServerConfig   `json:"server"`
	Jwt      JwtConfig      `json:"jwt"`
	Admin    AdminConfig    `json:"admin"`
	Postgres PostgresConfig `json:"postgresql"`
}

type ServerConfig struct {
	Port      string `json:"port"`
	ApiPrefix string `json:"api_prefix"`
}

type JwtConfig struct {
	Secret string `json:"secret"`
	Exp    int64  `json:"exp"`
}

type AdminConfig struct {
	AdminName string `json:"adminName"`
	Password  string `json:"adminPass"`
}

type PostgresConfig struct {
	Dsn string `json:"dsn"`
}

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
