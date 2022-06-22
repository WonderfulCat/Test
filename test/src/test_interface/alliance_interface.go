package test_interface

type AllianceI interface {
	Build(name string)
	AllianceOpeartionI
	AllianceStoreOpeartionI
}

//公会基础操作
type AllianceOpeartionI interface {
	CheckPermission(cname string, permission int32) bool
	CreateAlliance(aname string) bool
	JoinAlliance(cname string, permission int32) (bool, error)
	DismissAlliance(aname string) (bool, error)
	GetMemberList() string
	GetName() string
}

//公会仓库操作
type AllianceStoreOpeartionI interface {
	IncreaseCapacity() (bool, error)
	StoreItem(itemId, itemNum, index int32) (bool, error)
	DestoryItem(index int32) (bool, error)
	GetStoreList() string
	ClearUp()
}
