package Client

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/Jordanzuo/goutil/debugUtil"
	"github.com/Jordanzuo/goutil/logUtil"
)

//Start SocketClient开始入口
func Start(wg *sync.WaitGroup, socketServerAddress string) {
	defer func() {
		wg.Done()
	}()

	conn, err := net.DialTimeout("tcp", socketServerAddress, time.Second*2)
	if err != nil {
		logUtil.Log(fmt.Sprintf("连接服务器出错:%s", err.Error()), logUtil.Error, true)
		return
	}

	debugUtil.Println(fmt.Sprintf("client connection success, addr : %s", socketServerAddress))

	go handleRecData(conn)

	handleSendData(conn)
}
