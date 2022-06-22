package net_impl

import "test/src/test_net/net_interface"

//实现router时，先嵌入这个基类，然后根据需要对这个基类的方法进行重写
type BaseRouter struct{}

//这里之所以BaseRouter的方法都为空，
// 是因为有的Router不希望有PreHandle或PostHandle
// 所以Router全部继承BaseRouter的好处是，不需要实现PreHandle和PostHandle也可以实例化
func (br *BaseRouter) PreHandle(req net_interface.RequestI)  {}
func (br *BaseRouter) Handle(req net_interface.RequestI)     {}
func (br *BaseRouter) PostHandle(req net_interface.RequestI) {}
