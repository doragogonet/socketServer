package Server

import (
	"fmt"
	"net"
	"sync"

	"github.com/Jordanzuo/goutil/debugUtil"
	"github.com/Jordanzuo/goutil/logUtil"
)

//Start SocketServer启动入口
func Start(wg *sync.WaitGroup, socketServerAddress string) {
	debugUtil.Println("SocketServer Start")

	defer func() {
		wg.Done()
	}()

	listener, err := net.Listen("tcp", socketServerAddress)
	if err != nil {
		logUtil.Log(fmt.Sprintf("监听客户端出错:%s", err.Error()), logUtil.Error, true)
		return
	}

	debugUtil.Println(fmt.Sprintf("SocketServer Start, listen %s", listener.Addr()))

	for {
		// 阻塞直到新连接到来
		conn, err := listener.Accept()
		if err != nil {
			logUtil.Log(fmt.Sprintf("接收数据出错:%s", err.Error()), logUtil.Error, true)
			return
		}

		go handleReadData(conn)
	}
}
