package consul

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"usersvr/log"
)

func TestInit(t *testing.T) {
	register := NewRegister("127.0.0.1", 8500)

	err := register.Register("127.0.0.1", 8002, "usersvr", []string{""}, "aa8d6sa465465d3sad")

	if err != nil {
		fmt.Println(err)
	}

	quit := make(chan os.Signal)
	//告诉操作系统，当接收到 SIGINT 或 SIGTERM 这个信号时候会将信号传给通道
	//这样做的好处是可以在操作系统接收到 中断或者关闭时候优雅的执行后续程序
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//等待接收信号，阻塞
	<-quit
	// 服务终止，注销 consul 服务
	if err = register.DeRegister("7adsa5465a1ds3ad"); err != nil {
		log.Info("注销失败")
	} else {
		log.Info("注销成功")
	}
}
