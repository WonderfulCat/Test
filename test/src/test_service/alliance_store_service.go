package test_service

import (
	"encoding/json"
	"fmt"
	"test/src/test_common"
	"test/src/test_constant"
	"test/src/test_model"
	"test/src/test_net/net_impl"
	"test/src/test_net/net_interface"
)

//----------------------------------------------------------IncreaseCapacityHandle-------------------------------------------------//
type IncreaseCapacityHandle struct {
	net_impl.BaseRouter
}

func (c *IncreaseCapacityHandle) Handle(req net_interface.RequestI) {
	ret := c.IncreaseCapacity(req)
	if err := req.GetConnection().SendBuffMsg(req.GetMsgID(), GetJsonBytes(ret)); err != nil {
		fmt.Println(err)
	}
}

//仓库扩容
func (c *IncreaseCapacityHandle) IncreaseCapacity(request net_interface.RequestI) *test_model.ResponseInfo {
	character := test_common.CacheMap.CharacterGetByName(request.GetConnection().GetName())
	if character == nil {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_17}
	}

	if len(character.GetAllianceName()) <= 0 {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_14}
	}

	alliance, ok := test_common.CacheMap.AllianceGetByName(character.GetAllianceName())
	if !ok {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_13}
	}

	if !alliance.CheckPermission(request.GetConnection().GetName(), test_constant.MEMBER_PERMISSION_ADMIN) {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_15}
	}

	ok, err := alliance.IncreaseCapacity()
	if !ok {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: err.Error()}
	}
	return &test_model.ResponseInfo{Code: test_constant.RES_OK, Msg: err.Error()}
}

//----------------------------------------------------------StoreItemHandle-------------------------------------------------//
type StoreItemHandle struct {
	net_impl.BaseRouter
}

func (c *StoreItemHandle) Handle(req net_interface.RequestI) {
	ret := c.StoreItem(req)
	if err := req.GetConnection().SendBuffMsg(req.GetMsgID(), GetJsonBytes(ret)); err != nil {
		fmt.Println(err)
	}
}

//存储物品
func (c *StoreItemHandle) StoreItem(request net_interface.RequestI) *test_model.ResponseInfo {
	info := &test_model.StoreItemRequestInfo{}
	if err := json.Unmarshal(request.GetData(), info); err != nil {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: fmt.Sprintf(test_constant.RES_ERR_MSG_16, err.Error())}
	}

	character := test_common.CacheMap.CharacterGetByName(request.GetConnection().GetName())
	if character == nil {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_17}
	}

	if len(character.GetAllianceName()) <= 0 {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_14}
	}

	alliance, ok := test_common.CacheMap.AllianceGetByName(character.GetAllianceName())
	if !ok {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_13}
	}

	ok, err := alliance.StoreItem(info.ItemId, info.ItemNum, info.Index)
	if !ok {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: err.Error()}
	}
	return &test_model.ResponseInfo{Code: test_constant.RES_OK, Msg: err.Error()}
}

//----------------------------------------------------------DestoryItemHandle-------------------------------------------------//
type DestoryItemHandle struct {
	net_impl.BaseRouter
}

func (c *DestoryItemHandle) Handle(req net_interface.RequestI) {
	ret := c.DestoryItem(req)
	if err := req.GetConnection().SendBuffMsg(req.GetMsgID(), GetJsonBytes(ret)); err != nil {
		fmt.Println(err)
	}
}

//删除物品
func (c *DestoryItemHandle) DestoryItem(request net_interface.RequestI) *test_model.ResponseInfo {
	info := &test_model.DestoryItemRequestInfo{}
	if err := json.Unmarshal(request.GetData(), info); err != nil {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: fmt.Sprintf(test_constant.RES_ERR_MSG_16, err.Error())}
	}

	character := test_common.CacheMap.CharacterGetByName(request.GetConnection().GetName())
	if character == nil {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_17}
	}

	if len(character.GetAllianceName()) <= 0 {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_14}
	}

	alliance, ok := test_common.CacheMap.AllianceGetByName(character.GetAllianceName())
	if !ok {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_13}
	}

	if !alliance.CheckPermission(request.GetConnection().GetName(), test_constant.MEMBER_PERMISSION_ADMIN) {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_15}
	}

	ok, err := alliance.DestoryItem(info.Index)
	if !ok {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: err.Error()}
	}
	return &test_model.ResponseInfo{Code: test_constant.RES_OK, Msg: err.Error()}
}

//----------------------------------------------------------ClearUpHandle-------------------------------------------------//
type ClearUpHandle struct {
	net_impl.BaseRouter
}

func (c *ClearUpHandle) Handle(req net_interface.RequestI) {
	ret := c.ClearUp(req)
	if err := req.GetConnection().SendBuffMsg(req.GetMsgID(), GetJsonBytes(ret)); err != nil {
		fmt.Println(err)
	}
}

//整理物品
func (c *ClearUpHandle) ClearUp(request net_interface.RequestI) *test_model.ResponseInfo {
	character := test_common.CacheMap.CharacterGetByName(request.GetConnection().GetName())
	if character == nil {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_17}
	}

	if len(character.GetAllianceName()) <= 0 {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_14}
	}

	alliance, ok := test_common.CacheMap.AllianceGetByName(character.GetAllianceName())
	if !ok {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_13}
	}

	if !alliance.CheckPermission(request.GetConnection().GetName(), test_constant.MEMBER_PERMISSION_ADMIN) {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_15}
	}

	alliance.ClearUp()

	return &test_model.ResponseInfo{Code: test_constant.RES_OK, Msg: fmt.Sprintf(test_constant.RES_ERR_MSG_23, alliance.GetStoreList())}
}

//----------------------------------------------------------GetItemListHandle-------------------------------------------------//
type GetItemListHandle struct {
	net_impl.BaseRouter
}

func (c *GetItemListHandle) Handle(req net_interface.RequestI) {
	ret := c.GetItemList(req)
	if err := req.GetConnection().SendBuffMsg(req.GetMsgID(), GetJsonBytes(ret)); err != nil {
		fmt.Println(err)
	}
}

//整理物品
func (c *GetItemListHandle) GetItemList(request net_interface.RequestI) *test_model.ResponseInfo {
	character := test_common.CacheMap.CharacterGetByName(request.GetConnection().GetName())
	if character == nil {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_17}
	}

	if len(character.GetAllianceName()) <= 0 {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_14}
	}

	alliance, ok := test_common.CacheMap.AllianceGetByName(character.GetAllianceName())
	if !ok {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_13}
	}

	if !alliance.CheckPermission(request.GetConnection().GetName(), test_constant.MEMBER_PERMISSION_ADMIN) {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_15}
	}

	return &test_model.ResponseInfo{Code: test_constant.RES_OK, Msg: fmt.Sprintf(test_constant.RES_ERR_MSG_23, alliance.GetStoreList())}
}
