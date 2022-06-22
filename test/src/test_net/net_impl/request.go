package net_impl

import "test/src/test_net/net_interface"

type Request struct {
	conn net_interface.ConnectionI //已经和客户端建立好的 链接
	msg  net_interface.MessageI    //客户端请求的数据
}

func NewRequest(conn net_interface.ConnectionI, msg net_interface.MessageI) net_interface.RequestI {
	return &Request{conn: conn, msg: msg}
}

//获取请求连接信息
func (r *Request) GetConnection() net_interface.ConnectionI {
	return r.conn
}

//获取请求消息的数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

//获取请求的消息的ID
func (r *Request) GetMsgID() int32 {
	return r.msg.GetMsgId()
}
