package core

import (
	"GameManageSystem/global"
	"bytes"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

// 定义颜色
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

// InitLogger 初始化日志为自定义的一些类型
func InitLogger() *logrus.Logger {
	// 新建一个logger实例
	newLog := logrus.New()
	// 设置输出类型
	newLog.SetOutput(os.Stdout)
	// 开启返回函数名和行号
	newLog.SetReportCaller(global.Config.Logger.ShowLine)
	// 设置自己定义的Formatter
	newLog.SetFormatter(&LogFormatter{})
	// 从配置中获取日志级别并解析成 logrus 中的 Level
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		// 获取日志级别出错就设置日志级别为Info
		level = logrus.InfoLevel
	}
	//设置日志记录器的最低输出级别为上一步解析得到的级别
	newLog.SetLevel(level)
	// 设置全局默认日志类型
	InitDefaultLogger()
	log.Println("日志初始化成功")
	return newLog
}

// InitDefaultLogger 初始化全局log——logrus 以便后续mysql使用
func InitDefaultLogger() {
	//设置输出类型
	logrus.SetOutput(os.Stdout)
	//开启返回函数名和行号
	logrus.SetReportCaller(global.Config.Logger.ShowLine)
	//设置自己定义的Formatter
	logrus.SetFormatter(&LogFormatter{})
	//从配置中获取日志级别并解析成 logrus 中的 Level
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	//设置日志记录器的最低输出级别为上一步解析得到的级别
	logrus.SetLevel(level)
}

// Format 重写Format方法，自定义自己的日志类型
// 实现Formatter(entry *logrus.Entry) ([]byte, error)接口
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 根据不同的level去展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	// 初始化buffer 用于存储日志信息
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	log := global.Config.Logger

	// 自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")

	// entry.HasCaller() 判断日志条目中是否有调用者信息
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		_, _ = fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", log.Prefix, timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		_, _ = fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s\n", log.Prefix, timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}
