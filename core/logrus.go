package core

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"os"
	"path"
)

// 定义颜色常量
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

// Format 实现 Formatter 接口
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 根据不同的 Level 去展示颜色
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

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	log := global.Config.Logger
	// 自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")

	if entry.HasCaller() {
		// 获取文件路径和函数名
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		// 自定义输出格式
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", log.Prefix, timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		// 不带 Caller 的日志输出
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s\n", timestamp, levelColor, entry.Level, entry.Message)
	}

	return b.Bytes(), nil
}

// InitLogger 初始化自定义日志
func InitLogger() *logrus.Logger {
	mLog := logrus.New()
	// 设置输出到标准输出
	mLog.SetOutput(os.Stdout)
	// 开启返回函数名和行号
	//mLog.SetReportCaller(true)
	mLog.SetReportCaller(global.Config.Logger.ShowLine)
	// 设置自定义的 Formatter
	mLog.SetFormatter(&LogFormatter{})
	// 设置最低的日志级别
	//mLog.SetLevel(logrus.DebugLevel)
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	mLog.SetLevel(level)
	InitDefaultLogger()
	return mLog
}

// InitDefaultLogger 初始化默认全局日志
func InitDefaultLogger() {
	// 设置全局日志的输出类型
	logrus.SetOutput(os.Stdout)
	// 开启返回函数名和行号
	//logrus.SetReportCaller(true)
	logrus.SetReportCaller(global.Config.Logger.ShowLine)
	// 设置自定义的 Formatter
	logrus.SetFormatter(&LogFormatter{})
	// 设置最低的日志级别
	//logrus.SetLevel(logrus.DebugLevel)
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
}
