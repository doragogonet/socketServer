package main

import (
	"sync"

	"wshhz.com/Socket/SocketServer/SocketClient/config"
	"wshhz.com/Socket/SocketServer/SocketClient/src/Client"
	"wshhz.com/Socket/SocketServer/SocketClient/src/signalMgr"
)

var (
	wg sync.WaitGroup
)

func init() {
	wg.Add(1)
}

func main() {
	signalMgr.Start(nil, nil)

	Client.Start(&wg, config.SocketServerAddress)

	// 防止主线程退出
	wg.Wait()
}
