package Server

import (
	"fmt"
	"net"
	"sync"
)

//Start SocketServer启动入口
func Start(wg *sync.WaitGroup, socketServerAddress string) {
	fmt.Println("SocketServer Start")

	defer func() {
		wg.Done()
	}()

	listener, err := net.Listen("tcp", socketServerAddress)
	if err != nil {
		fmt.Println("listen err : " + err.Error())
		return
	}

	fmt.Println(fmt.Sprintf("SocketServer Start, listen %s", listener.Addr()))

	for {
		// 阻塞直到新连接到来
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error : " + err.Error())
			return
		}

		go handleReadData(conn)
	}
}
