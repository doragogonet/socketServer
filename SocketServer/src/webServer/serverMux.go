package webServer

import (
	"net/http"

	"github.com/Jordanzuo/goutil/logUtil"
)

// 自定义路由对象
type selfDefineMux struct {
}

//ServeHTTP 实现默认路由的自定义方法
func (mux *selfDefineMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 获取对应的handle对象
	handleObj, exist := getHandle(r.RequestURI)
	if !exist {
		logUtil.Log("API不存在:"+r.RequestURI, logUtil.Error, true)
		return
	}

	// 调用注册的回调函数
	result := handleObj.HandleFunc()

	responseResultDefault(w, result)
}
