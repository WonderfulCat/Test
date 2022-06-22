package net_impl

import (
	"test/src/test_net/net_interface"
)

type MsgHandle struct {
	Routers        map[int32]net_interface.RouterI //存放每个MsgId 所对应的处理方法的map属性
	WorkerPoolSize uint32                          //业务工作Worker池的数量
	TaskQueue      chan net_interface.RequestI     //Worker负责取任务的消息队列
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		WorkerPoolSize: 1,
		Routers:        make(map[int32]net_interface.RouterI),
		TaskQueue:      make(chan net_interface.RequestI, 4096),
	}
}

//将消息交给TaskQueue,由worker进行处理
func (c *MsgHandle) SendMsgToTaskQueue(request net_interface.RequestI) {
	//将请求消息发送给任务队列
	c.TaskQueue <- request
}

//启动一个Worker工作流程
func (c *MsgHandle) StartWorker() {
	//不断的等待队列中的消息
	for request := range c.TaskQueue {
		c.DoMsgHandler(request)
	}
}

//马上以非阻塞方式处理消息
func (mh *MsgHandle) DoMsgHandler(request net_interface.RequestI) {
	handler, ok := mh.Routers[request.GetMsgID()]
	if !ok {
		return
	}

	//执行对应处理方法
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

//为消息添加具体的处理逻辑
func (mh *MsgHandle) AddRouter(msgId int32, router net_interface.RouterI) {
	if _, ok := mh.Routers[msgId]; ok {
		return
	}

	mh.Routers[msgId] = router
}
