package webServer

import (
	"encoding/json"

	"github.com/Jordanzuo/goutil/logUtil"
	"wshhz.com/Socket/SocketServer/SocketServer/src/Server"
	. "wshhz.com/Socket/SocketServer/SocketServer/src/model"
)

func init() {
	regesiterHandle("/API/getClientList", getClientList)
}

func getClientList() *Result {
	clientList := Server.GetClientList()

	result, err := json.Marshal(clientList)
	if err != nil {
		logUtil.Log("json解析数据错误:"+err.Error(), logUtil.Error, true)
	}

	return NewResult(string(result))
}
