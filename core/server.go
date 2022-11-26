package core

import (
	"fmt"
	"go.uber.org/zap"
	"singer/global"
	"singer/initialize"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {

	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.Config.System.Addr)
	s := initServer(address, Router)

	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.Logger.Info("server run success on ", zap.String("address", address))

	fmt.Printf("singer is running... \n", address)
	global.Logger.Error(s.ListenAndServe().Error())
}
