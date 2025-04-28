package main

import (
	"go.uber.org/zap"
	"os"
	"os/signal"
	"ranbao/ini"
	"ranbao/web/handler"
	"syscall"
)

func main() {
	ini.InitConfig("./config/config.yaml")

	ini.InitLog()
	ini.InitDB()
	ini.InitRDB()

	r := handler.InitRouter()
	//启动服务
	go func() {
		zap.S().Debugf("服务启动")
		if err := r.Run("127.0.0.1:8000"); err != nil {
			zap.S().Panic("启动失败", err.Error())
		}
	}()

	//接收终止信号 Signal表示操作系统信号
	qiut := make(chan os.Signal, 1)
	//接收control+c
	signal.Notify(qiut, syscall.SIGINT, syscall.SIGTERM)
	<-qiut
}
