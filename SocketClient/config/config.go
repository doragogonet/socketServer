package config

import (
	"github.com/Jordanzuo/goutil/configUtil"
	"github.com/Jordanzuo/goutil/debugUtil"
	"github.com/Jordanzuo/goutil/logUtil"
)

var (
	//DEBUG  是否是DEBUG模式
	DEBUG bool

	//SocketServerAddress SocketServer监听地址
	SocketServerAddress string
)

func init() {
	// 设置日志文件的存储目录
	logUtil.SetLogPath("LOG")

	// 读取配置文件内容
	config, err := configUtil.ReadJsonConfig("config.ini")
	checkError(err)

	// 解析DEBUG配置
	debug, err := configUtil.ReadBoolJsonValue(config, "DEBUG")
	checkError(err)

	// 为DEBUG模式赋值
	DEBUG = debug

	// 设置debugUtil的状态
	debugUtil.SetDebug(debug)

	// 解析SocketServer监听地址配置数据
	SocketServerAddress, err = configUtil.ReadStringJsonValue(config, "SocketServerAddress")
	checkError(err)

	debugUtil.Println("DEBUG:", debug)
	debugUtil.Println("SocketServerAddress", SocketServerAddress)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
