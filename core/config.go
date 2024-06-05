package core

import (
	"GameManageSystem/config"
	"GameManageSystem/global"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/fs"
	"log"
	"os"
)

const ConfigFile = "./settings.yaml"

// InitConfig 读取yaml配置文件
func InitConfig() {
	//使用地址，防止内存释放
	c := &config.Config{}
	// 读取文件 读取出来的内容为字节类型的切片
	yamlConfigFile, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error:%s", err))
	}
	// 反序列化将配置信息转换成结构体里面的信息储存起来
	err = yaml.Unmarshal(yamlConfigFile, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	// 读取配置成功，记录日志
	log.Println("配置信息加载成功")
	// 将读取好的配置给全局变量
	global.Config = c
}

func SetYaml() error {
	// 序列化，将结构体里面的数据转化为字节类型的切片
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	// 写入文件，把转化的内容写入到对应文件当中去WriteFile(写入哪个文件，写入的内容，权限)
	err = os.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		return err
	}
	// 写入成功，记录日志
	log.Println("配置信息修改成功")
	return nil
}
