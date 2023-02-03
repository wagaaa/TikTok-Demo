package utils

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config 读取config.yaml并绑定至cfg结构体
type Config struct {
	Database struct {
		User   string
		Passwd string
		Host   string
		Port   string
		Name   string
	}
}

var Cfg Config

func init() {
	//var cfg Config
	buf, err := os.ReadFile("biz/config/config.yaml")
	if err != nil {
		println(err)
	}
	err = yaml.Unmarshal(buf, &Cfg)
	if err != nil {
		println(err)
	}
}
