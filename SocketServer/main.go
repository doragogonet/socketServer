package main

import (
	"sync"
	"time"

	"wshhz.com/Socket/SocketServer/SocketServer/config"
	"wshhz.com/Socket/SocketServer/SocketServer/src/Server"
	"wshhz.com/Socket/SocketServer/SocketServer/src/signalMgr"
)

var (
	wg sync.WaitGroup
)

func init() {
	wg.Add(1)
}

func main() {
	signalMgr.Start(nil, nil)

	go Server.Start(&wg, config.SocketServerAddress)

	go testSend()

	// 防止主线程退出
	wg.Wait()
}

// 测试发送数据
func testSend() {
	defer wg.Done()

	testData := "testSendData"
	for {
		// 每秒向所有注册的客户端发送一次数据
		Server.SendAllClientData([]byte(testData))
		time.Sleep(time.Second * 1)
	}
}
