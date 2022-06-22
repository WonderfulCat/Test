package net_interface

//定义服务器接口
type ServerI interface {
	//启动服务器方法
	Start()
	//停止服务器方法
	Stop()
	//路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
	AddRouter(msgId int32, router RouterI)
	//设置该Server的连接创建时Hook函数
	SetOnConnStart(func(ConnectionI))
	//设置该Server的连接断开时的Hook函数
	SetOnConnStop(func(ConnectionI))
	//调用连接OnConnStart Hook函数
	CallOnConnStart(conn ConnectionI)
	//调用连接OnConnStop Hook函数
	CallOnConnStop(conn ConnectionI)
}
