package model

//Handle api对应的处理对象
type Handle struct {
	path       string
	HandleFunc func() *Result
}

//NewHandle 新建一个处理对象
// 参数:
// _path:路径
// handleFunc:对应的回调函数
func NewHandle(_path string, handleFunc func() *Result) *Handle {
	return &Handle{
		path:       _path,
		HandleFunc: handleFunc,
	}
}
