package test_service

import (
	"encoding/json"
	"fmt"

	"test/src/test_common"
	"test/src/test_constant"
	"test/src/test_interface"

	"test/src/test_model"
	"test/src/test_net/net_impl"
	"test/src/test_net/net_interface"
)

//----------------------------------------------------------WhichAllianceHandle-------------------------------------------------//
type WhichAllianceHandle struct {
	net_impl.BaseRouter
}

func (c *WhichAllianceHandle) Handle(req net_interface.RequestI) {
	ret := c.WhichAlliance(req)
	if err := req.GetConnection().SendBuffMsg(req.GetMsgID(), GetJsonBytes(ret)); err != nil {
		fmt.Println(err)
	}
}

//查询自己所在公会 || 公会列表
func (c *WhichAllianceHandle) WhichAlliance(request net_interface.RequestI) *test_model.ResponseInfo {
	character := test_common.CacheMap.CharacterGetByName(request.GetConnection().GetName())
	if character == nil {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_17}
	}

	if len(character.GetAllianceName()) <= 0 {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_4}
	}

	//所有公会名称列表
	alliance, ok := test_common.CacheMap.AllianceGetByName(character.GetAllianceName())
	if !ok || alliance == nil {
		character.SetAllianceName("")
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_4}
	}

	return &test_model.ResponseInfo{Code: test_constant.RES_OK, Msg: alliance.GetMemberList()}
}

//----------------------------------------------------------AllianceListHandle-------------------------------------------------//

type AllianceListHandle struct {
	net_impl.BaseRouter
}

func (c *AllianceListHandle) Handle(req net_interface.RequestI) {
	ret := c.AllianceList(req)
	if err := req.GetConnection().SendBuffMsg(req.GetMsgID(), GetJsonBytes(ret)); err != nil {
		fmt.Println(err)
	}
}

//公会列表
func (c *AllianceListHandle) AllianceList(request net_interface.RequestI) *test_model.ResponseInfo {
	return &test_model.ResponseInfo{Code: test_constant.RES_OK, Msg: fmt.Sprintf(test_constant.RES_ERR_MSG_19, test_common.CacheMap.AllianceList())}
}

//----------------------------------------------------------CreateAllianceHandle-------------------------------------------------//
type CreateAllianceHandle struct {
	net_impl.BaseRouter
}

func (c *CreateAllianceHandle) Handle(req net_interface.RequestI) {
	ret := c.CreateAlliance(req)
	if err := req.GetConnection().SendBuffMsg(req.GetMsgID(), GetJsonBytes(ret)); err != nil {
		fmt.Println(err)
	}
}

//创建公会 cname:角色名称  aname:公会名称
func (c *CreateAllianceHandle) CreateAlliance(request net_interface.RequestI) *test_model.ResponseInfo {
	info := &test_model.CreateAllianceRequestInfo{}
	if err := json.Unmarshal(request.GetData(), info); err != nil {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: fmt.Sprintf(test_constant.RES_ERR_MSG_16, err.Error())}
	}

	character := test_common.CacheMap.CharacterGetByName(request.GetConnection().GetName())
	if character == nil {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_17}
	}

	if len(character.GetAllianceName()) > 0 {
		_, ok := test_common.CacheMap.AllianceGetByName(character.GetAllianceName())
		if ok {
			return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_11}
		}
	}

	//创建
	alliance := GetReflectByName(test_constant.REGISTER_NAME_ALLIANCE).(test_interface.AllianceI)
	alliance.Build(info.AName)
	//设置权限
	alliance.JoinAlliance(character.GetName(), test_constant.MEMBER_PERMISSION_ADMIN)
	//缓存
	if !test_common.CacheMap.AllianceAdd(alliance) {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_12}
	}
	//更新角色
	character.SetAllianceName(info.AName)

	return &test_model.ResponseInfo{Code: test_constant.RES_OK, Msg: fmt.Sprintf(test_constant.RES_ERR_MSG_20, info.AName)}
}

//----------------------------------------------------------JoinAllianceHandle-------------------------------------------------//
type JoinAllianceHandle struct {
	net_impl.BaseRouter
}

func (c *JoinAllianceHandle) Handle(req net_interface.RequestI) {
	ret := c.JoinAlliance(req)
	if err := req.GetConnection().SendBuffMsg(req.GetMsgID(), GetJsonBytes(ret)); err != nil {
		fmt.Println(err)
	}
}

//加入公会 cname:角色名称  aname:公会名称
func (c *JoinAllianceHandle) JoinAlliance(request net_interface.RequestI) *test_model.ResponseInfo {
	info := &test_model.JoinAllianceRequestInfo{}
	if err := json.Unmarshal(request.GetData(), info); err != nil {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: fmt.Sprintf(test_constant.RES_ERR_MSG_16, err.Error())}
	}

	character := test_common.CacheMap.CharacterGetByName(request.GetConnection().GetName())
	if character == nil {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_17}
	}

	if len(character.GetAllianceName()) > 0 {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_10}
	}

	alliance, ok := test_common.CacheMap.AllianceGetByName(info.AName)
	if !ok {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_13}
	}

	if ok, err := alliance.JoinAlliance(info.AName, test_constant.MEMBER_PERMISSION_NORMAL); !ok {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: err.Error()}
	}

	return &test_model.ResponseInfo{Code: test_constant.RES_OK}
}

//----------------------------------------------------------DismissAllianceHandle-------------------------------------------------//
type DismissAllianceHandle struct {
	net_impl.BaseRouter
}

func (c *DismissAllianceHandle) Handle(req net_interface.RequestI) {
	ret := c.DismissAlliance(req)
	if err := req.GetConnection().SendBuffMsg(req.GetMsgID(), GetJsonBytes(ret)); err != nil {
		fmt.Println(err)
	}
}

//解散公会
func (c *DismissAllianceHandle) DismissAlliance(request net_interface.RequestI) *test_model.ResponseInfo {
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

	if ok, err := alliance.DismissAlliance(request.GetConnection().GetName()); !ok {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: err.Error()}
	}

	test_common.CacheMap.AllianceRemove(alliance.GetName())
	return &test_model.ResponseInfo{Code: test_constant.RES_OK, Msg: fmt.Sprintf(test_constant.RES_ERR_MSG_21, alliance.GetName())}
}
