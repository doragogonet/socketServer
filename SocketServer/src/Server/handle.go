package Server

import (
	"fmt"
	"net"
	"time"

	"github.com/Jordanzuo/goutil/debugUtil"
	"github.com/Jordanzuo/goutil/logUtil"
)

// 处理客户端连接
func handleReadData(conn net.Conn) {
	clientObj := newClientObj(conn)

	registerClient(clientObj)
	defer func() {
		unRegisterClient(clientObj)
	}()

	// 启动发送数据的协程
	go handleSendData(clientObj)

	for {
		readBytes := make([]byte, 1024)

		// Read方法会阻塞直到收到数据
		n, err := conn.Read(readBytes)
		if err != nil {
			logUtil.Log(fmt.Sprintf("读取客户端数据出错:%s", err.Error()), logUtil.Error, true)
			return
		}

		// 添加数据
		clientObj.appendRecData(readBytes[:n])

		handleRecData(clientObj)
	}
}

// 处理接收到客户端的数据
// client:客户端对象
func handleRecData(client *clientObj) {
	// 接收到的消息长度
	n := len(client.recData)

	// 转换成字符串
	message := string(client.recData[:n])

	// 截取数据,得到新的数据
	client.recData = client.recData[n:]

	debugUtil.Println(fmt.Sprintf("收到来自id=%d的客户端的消息:%s", client.id, message))
}

// 处理客户端发送的数据
// client:客户端对象
func handleSendData(client *clientObj) {
	for {
		message, exist := client.getSendData()
		if !exist {
			// 不存在要发送的数据,休眠5毫秒
			time.Sleep(time.Millisecond * 5)
			continue
		}

		// 发送数据
		n, err := client.conn.Write(message.data)
		if err != nil {
			logUtil.Log(fmt.Sprintf("本次发送数据错误:%s", err.Error()), logUtil.Error, true)
			continue
		}

		debugUtil.Println(fmt.Sprintf("本次发送给id=%d的客户端长度为%d的数据", client.id, n))
	}
}
