package Server

import (
	"sync/atomic"
)

var (
	// 全局自增的消息对象Id
	globalDataId int32
)

// 发送的信息对象
type sendDataItem struct {
	// id
	id int32

	// 数据
	data []byte
}

// 新建一个发送的消息对象
func newSendDataItem(_data []byte) *sendDataItem {
	return &sendDataItem{
		id:   atomic.AddInt32(&globalDataId, 1),
		data: _data,
	}
}
