package main

import (
	"fmt"
	"sync"
	"time"

	"wshhz.com/Socket/SocketServer/SocketServer/Server"
)

var (
	wg sync.WaitGroup
)

const (
	socketServerAddress = "10.254.0.129:8090"
)

func init() {
	wg.Add(1)
}

func main() {
	go Server.Start(&wg, socketServerAddress)

	go testSend()

	wg.Wait()

	fmt.Println("SocketServer End...")
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
