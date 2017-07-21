package Client

import (
	"fmt"
	"net"
	"time"
)

var (
	testString = "testClient"
)

// 处理发送数据
func handleSendData(conn net.Conn) {
	for {
		// fmt.Print("please input send data:")
		// fmt.Scanln(&testString)

		n, err := conn.Write([]byte(testString))
		if err != nil {
			fmt.Println("send err :" + err.Error())
			return
		}

		fmt.Println(fmt.Sprintf("send data len:%d", n))

		// 测试通信,每秒发送一次
		time.Sleep(time.Second * 1)
	}
}

// 处理收到的数据
func handleRecData(conn net.Conn) {
	for {
		recData := make([]byte, 1024)
		n, err := conn.Read(recData)
		if err != nil {
			fmt.Println("read err:" + err.Error())
		}

		message := string(recData[:n])
		fmt.Println(fmt.Sprintf("收到来自%s的消息:%s", conn.RemoteAddr().String(), message))
	}
}
