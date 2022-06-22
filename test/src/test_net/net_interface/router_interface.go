package net_interface

type RouterI interface {
	PreHandle(request RequestI)  //在处理conn业务之前的钩子方法
	Handle(request RequestI)     //处理conn业务的方法
	PostHandle(request RequestI) //处理conn业务之后的钩子方法
}
