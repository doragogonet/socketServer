package webServer

import (
	"net/http"
	"sync"

	"github.com/Jordanzuo/goutil/logUtil"
)

//Start webServer开始启动入口
func Start(wg *sync.WaitGroup, webServerAddress string) {
	defer func() {
		wg.Done()
	}()

	logUtil.Log("webServer start listen:"+webServerAddress, logUtil.Info, true)

	// 开始监听webServer
	if err := http.ListenAndServe(webServerAddress, new(selfDefineMux)); err != nil {
		logUtil.Log("webServer Listen Error:"+err.Error(), logUtil.Error, true)
	}
}
