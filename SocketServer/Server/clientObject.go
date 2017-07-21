package Server

import (
	"net"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// 全局自增的客户端Id
	globalId int32
)

type clientObj struct {
	// 客户端连接对象
	conn net.Conn

	// id
	id int32

	// 接收到的消息
	recData []byte

	// 要发送的消息
	sendData []*sendDataItem

	// 上次活跃时间
	activeTime time.Time

	// 同步锁对象
	mutex sync.RWMutex
}

// 添加接收到的数据
func (obj *clientObj) appendRecData(_recData []byte) {
	obj.recData = append(obj.recData, _recData...)
	obj.activeTime = time.Now()
}

// 添加要发送的数据
func (obj *clientObj) appendSendData(_sendDataItem *sendDataItem) {
	obj.mutex.Lock()
	defer obj.mutex.Unlock()

	obj.sendData = append(obj.sendData, _sendDataItem)
}

// 获取要发送的数据
// 返回值:
// sendData:要发送的数据
// exist:是否存在要发送的数据
func (obj *clientObj) getSendData() (sendData *sendDataItem, exist bool) {
	obj.mutex.Lock()
	obj.mutex.Unlock()

	data := obj.sendData
	if len(data) == 0 {
		return nil, false
	}

	// 取出要发送的数据
	sendData = data[0]
	exist = true

	// 删除已取出的数据
	obj.sendData = obj.sendData[1:]

	return
}

// 新创建一个客户端对象
func newClientObj(_conn net.Conn) *clientObj {
	return &clientObj{
		conn:       _conn,
		id:         atomic.AddInt32(&globalId, 1),
		recData:    make([]byte, 0, 1024),
		sendData:   make([]*sendDataItem, 0, 16),
		activeTime: time.Now(),
	}
}
