package Server

import (
	"wshhz.com/Socket/SocketServer/SocketServer/src/model"
)

//GetClientList 获取客户端列表
// 返回值:
// []ClientObj:客户端列表
func GetClientList() []*model.ClientObj {
	// 获取当前注册的客户端列表
	clientList := getClientList()

	// 获取客户端的个数
	n := len(clientList)

	resultList := make([]*model.ClientObj, 0, n)

	for _, c := range clientList {
		resultList = append(resultList, model.NewClientObj(c.id, c.conn.RemoteAddr().String()))
	}

	return resultList
}
