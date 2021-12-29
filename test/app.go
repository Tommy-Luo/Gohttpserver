package main

type ServerConfig struct {

	Mysql mysql `yaml:"mysql"` // 嵌入MySQL配置
}