package config

import "fmt"

type System struct {
	Host string `yaml:"host"` // 主机ip
	Env  string `yaml:"env"`  // 版本
	Port int    `yaml:"port"` // 端口号
}

func (s System) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
