package main

import (
	"log"

	"github.com/yangtizi/go/config/jsonconfig"
)

// SConfig 配置
type SConfig struct {
	Redis string `json:"redis,omitempty"`
	MySQL string `json:"mysql,omitempty"`
}

var cfg *SConfig

func main() {
	loadConfigDemo()
}

func loadConfigDemo() {
	cfg = &SConfig{}
	jsonconfig.JSONParsing("./config.json", cfg)

	log.Println("加载的配置是", cfg)
}
