package config

import "strconv"

type Mysql struct {
	Host     string `yaml:"host"`      // ip地址
	Port     int    `yaml:"port"`      // 端口号
	Config   string `yaml:"config"`    // 高级配置 如：charset
	DB       string `yaml:"db"`        // 数据库名
	User     string `yaml:"user"`      // 用户名
	Password string `yaml:"password"`  // 用户密码
	Loglevel string `yaml:"log_level"` // 日志等级 debug就是输出全部sql，dev, release
}

func (m Mysql) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?" + m.Config
}
