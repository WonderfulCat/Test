package net_interface

type MsgHandleI interface {
	AddRouter(msgId int32, router RouterI) //为消息添加具体的处理逻辑
	StartWorker()                          //启动worker
	SendMsgToTaskQueue(request RequestI)   //将消息交给TaskQueue,由worker进行处理
}
