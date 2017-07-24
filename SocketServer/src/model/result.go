/*
结果对象文件,可自定义实现结果对象里的结构
*/

package model

//Result 结果对象
type Result struct {
	Data string
}

//NewResult 新建一个结果对象
// 参数:
// data:数据
// 返回值:
// *result:新建的结果对象
func NewResult(data string) *Result {
	return &Result{
		Data: data,
	}
}
