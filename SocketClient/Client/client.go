package Client

import (
	"fmt"
	"net"
	"sync"
	"time"
)

//Start SocketClient开始入口
func Start(wg *sync.WaitGroup, socketServerAddress string) {
	defer func() {
		wg.Done()
	}()

	conn, err := net.DialTimeout("tcp", socketServerAddress, time.Second*2)
	if err != nil {
		fmt.Println("connection err :" + err.Error())
		return
	}

	fmt.Println(fmt.Sprintf("client connection success, addr : %s", socketServerAddress))

	go handleRecData(conn)

	handleSendData(conn)
}
