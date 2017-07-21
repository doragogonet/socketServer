package Server

//SendAllClientData 发送消息给所有的客户端对象
// 参数:
// data:要发送的消息
func SendAllClientData(data []byte) {
	// 获取所有活跃的客户端列表
	clientList := getClientList()

	// 新建要发送的消息对象
	sendData := newSendDataItem(data)

	// 遍历所有的客户端对象
	for _, client := range clientList {
		// 发送消息
		client.appendSendData(sendData)
	}
}
