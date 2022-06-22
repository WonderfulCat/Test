package net_interface

type RequestI interface {
	GetConnection() ConnectionI //获取请求连接信息
	GetData() []byte            //获取请求消息的数据
	GetMsgID() int32            //获取请求的消息ID
}
