package webServer

import (
	"wshhz.com/Socket/SocketServer/SocketServer/src/Server"
	. "wshhz.com/Socket/SocketServer/SocketServer/src/model"
)

var (
	testString = "testSend"
)

func init() {
	regesiterHandle("/API/send", send)
}

func send() *Result {
	data := []byte(testString)
	Server.SendAllClientData(data)

	return NewResult("ok")
}
