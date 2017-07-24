package webServer

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jordanzuo/goutil/logUtil"
	"wshhz.com/Socket/SocketServer/SocketServer/src/model"
)

func responseResultDefault(w http.ResponseWriter, result *model.Result) {
	data, err := json.Marshal(result)
	if err != nil {
		logUtil.Log("json解析数据出错:"+err.Error(), logUtil.Error, true)
		return
	}

	n, err := w.Write(data)
	if err != nil {
		logUtil.Log("发送数据出错:"+err.Error(), logUtil.Error, true)
		return
	}

	logUtil.Log(fmt.Sprintf("本次发送长度为%d的数据", n), logUtil.Debug, true)
}
