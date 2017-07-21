package Server

import (
	"fmt"
	"sync"
)

var (
	// 客户端集合
	clientMap = make(map[int32]*clientObj)

	// 同步锁对象
	mutex sync.RWMutex
)

//  获取客户端列表
func getClientList() map[int32]*clientObj {
	mutex.RLock()
	defer mutex.RUnlock()

	return clientMap
}

// 注册客户端对象
func registerClient(obj *clientObj) {
	mutex.Lock()
	defer mutex.Unlock()

	clientMap[obj.id] = obj
	fmt.Println(fmt.Sprintf("注册客户端,id=%d,Addr=%s", obj.id, obj.conn.RemoteAddr().String()))
}

// 注销客户端对象
func unRegisterClient(obj *clientObj) {
	mutex.Lock()
	defer mutex.Unlock()

	delete(clientMap, obj.id)
	fmt.Println(fmt.Sprintf("注销客户端,id=%d", obj.id))
}
