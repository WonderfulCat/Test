package net_interface

import "net"

//定义连接接口
type ConnectionI interface {
	//启动连接，让当前连接开始工作
	Start()
	//停止连接，结束当前连接状态M
	Stop()
	//直接将数据发送给TCP客户端(有缓冲)
	SendBuffMsg(msgId int32, data []byte) error
	GetTCPConnection() *net.TCPConn
	SetName(name string)
	GetName() string
}
