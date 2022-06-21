package test_interface

import "test/src/test_model"

type AllianceI interface {
	Build(name string)
	AllianceOpeartionI
	AllianceStoreOpeartionI
}

//公会基础操作
type AllianceOpeartionI interface {
	CheckPermission(cname string, permission int32) bool
	CreateAlliance(aname string) *test_model.ResponseInfo
	JoinAlliance(cname string, permission int32) *test_model.ResponseInfo
	DismissAlliance(aname string) *test_model.ResponseInfo
	GetMemberList() string
	GetName() string
}

//公会仓库操作
type AllianceStoreOpeartionI interface {
	IncreaseCapacity() *test_model.ResponseInfo
	StoreItem(itemId, itemNum, index int32) *test_model.ResponseInfo
	DestoryItem(index int32) *test_model.ResponseInfo
	ClearUp()
}
