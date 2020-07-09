package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

// 配置日志输出
func init() {
	logrus.SetLevel(logrus.DebugLevel)

	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logfile, _ := os.OpenFile("./app.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	logrus.SetOutput(logfile) //默认为os.stderr
}

//方式一：logrus函数（最终调用的是logrus.StandardLogger默认实例方法）
func main() {
	logrus.Infoln("测试数据")
	logrus.Errorln("测试error")
}
