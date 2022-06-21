package test_net

import "test/src/test_model"

var Routers map[int32]func(conn *ConnTest, message *Message) *test_model.ResponseInfo

func init() {
	Routers = make(map[int32]func(conn *ConnTest, message *Message) *test_model.ResponseInfo)
}

func RegisterRouter(opCode int32, f func(conn *ConnTest, message *Message) *test_model.ResponseInfo) {
	Routers[opCode] = f
}
