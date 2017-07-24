package model

//ClientObj 客户端对象
type ClientObj struct {
	//Id id
	Id int32

	//Addr 地址
	Addr string
}

//NewClientObj 新建客户端对象
// 参数:
// id:id
// addr:地址
// 返回值:
// *ClientObj:客户端对象
func NewClientObj(id int32, addr string) *ClientObj {
	return &ClientObj{
		Id:   id,
		Addr: addr,
	}
}
