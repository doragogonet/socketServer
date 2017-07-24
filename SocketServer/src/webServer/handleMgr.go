/*
用来管理web请求的文件
*/

package webServer

import (
	"github.com/Jordanzuo/goutil/logUtil"
	. "wshhz.com/Socket/SocketServer/SocketServer/src/model"
)

var (
	// 保存handle对象的集合
	handleMap = make(map[string]*Handle)
)

// 注册api及对应的回调函数
// 参数:
// api:api地址
// callback:对应的回调函数
func regesiterHandle(api string, callback func() *Result) {
	if _, exist := handleMap[api]; exist {
		logUtil.LogAndPrint("已注册过该方法:"+api, logUtil.Error)
	}

	// 注册API及其对应的回调函数
	handleMap[api] = NewHandle(api, callback)
}

// 获取api对应的handle对象
// 参数:
// api:api地址
// 返回值:
// *handle:对应的handle对象
// bool:是否存在对应的handle对象
func getHandle(api string) (*Handle, bool) {
	handle, exist := handleMap[api]

	return handle, exist
}
