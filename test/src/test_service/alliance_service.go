package test_service

import (
	"encoding/json"
	"fmt"
	"test/src/test_common"
	"test/src/test_constant"
	"test/src/test_interface"
	"test/src/test_model"
	"test/src/test_net"
)

//查询自己所在公会 || 公会列表
func WhichAlliance(conn *test_net.ConnTest, message *test_net.Message) *test_model.ResponseInfo {
	character := test_common.CacheMap.CharacterGetByName(conn.Name)
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
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_9}
	}

	return &test_model.ResponseInfo{Code: test_constant.RES_OK, Msg: alliance.GetMemberList()}
}

//公会列表
func AllianceList(conn *test_net.ConnTest, message *test_net.Message) *test_model.ResponseInfo {
	return &test_model.ResponseInfo{Code: test_constant.RES_OK, Msg: test_common.CacheMap.AllianceList()}
}

//创建公会 cname:角色名称  aname:公会名称
func CreateAlliance(conn *test_net.ConnTest, message *test_net.Message) *test_model.ResponseInfo {
	info := &test_model.CreateAllianceRequestInfo{}
	if err := json.Unmarshal(message.GetData(), info); err != nil {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: fmt.Sprintf(test_constant.RES_ERR_MSG_16, err.Error())}
	}

	character := test_common.CacheMap.CharacterGetByName(conn.Name)
	if character == nil {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_17}
	}
	if len(character.GetAllianceName()) > 0 {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_11}
	}

	//创建
	alliance := test_common.GetReflectByName(test_constant.REGISTER_NAME_ALLIANCE).(test_interface.AllianceI)
	alliance.Build(info.AName)
	//设置权限
	alliance.JoinAlliance(character.GetName(), test_constant.MEMBER_PERMISSION_ADMIN)
	//缓存
	if !test_common.CacheMap.AllianceAdd(alliance) {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_12}
	}
	//更新角色
	character.SetAllianceName(info.AName)

	return &test_model.ResponseInfo{Code: test_constant.RES_OK}
}

//加入公会 cname:角色名称  aname:公会名称
func JoinAlliance(conn *test_net.ConnTest, message *test_net.Message) *test_model.ResponseInfo {
	info := &test_model.JoinAllianceRequestInfo{}
	if err := json.Unmarshal(message.GetData(), info); err != nil {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: fmt.Sprintf(test_constant.RES_ERR_MSG_16, err.Error())}
	}

	character := test_common.CacheMap.CharacterGetByName(conn.Name)
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

	return alliance.JoinAlliance(info.AName, test_constant.MEMBER_PERMISSION_NORMAL)
}

//解散公会
func DismissAlliance(conn *test_net.ConnTest, message *test_net.Message) *test_model.ResponseInfo {
	character := test_common.CacheMap.CharacterGetByName(conn.Name)
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

	res := alliance.DismissAlliance(conn.Name)
	if res.Code != test_constant.RES_OK {
		return res
	}

	test_common.CacheMap.AllianceRemove(alliance.GetName())
	return &test_model.ResponseInfo{Code: test_constant.RES_OK}
}

//仓库扩容
func IncreaseCapacity(conn *test_net.ConnTest, message *test_net.Message) *test_model.ResponseInfo {
	character := test_common.CacheMap.CharacterGetByName(conn.Name)
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

	if !alliance.CheckPermission(conn.Name, test_constant.MEMBER_PERMISSION_ADMIN) {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_15}
	}

	return alliance.IncreaseCapacity()
}

//存储物品
func StoreItem(conn *test_net.ConnTest, message *test_net.Message) *test_model.ResponseInfo {
	info := &test_model.StoreItemRequestInfo{}
	if err := json.Unmarshal(message.GetData(), info); err != nil {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: fmt.Sprintf(test_constant.RES_ERR_MSG_16, err.Error())}
	}

	character := test_common.CacheMap.CharacterGetByName(conn.Name)
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

	return alliance.StoreItem(info.ItemId, info.ItemNum, info.Index)
}

//删除物品
func DestoryItem(conn *test_net.ConnTest, message *test_net.Message) *test_model.ResponseInfo {
	info := &test_model.DestoryItemRequestInfo{}
	if err := json.Unmarshal(message.GetData(), info); err != nil {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: fmt.Sprintf(test_constant.RES_ERR_MSG_16, err.Error())}
	}

	character := test_common.CacheMap.CharacterGetByName(conn.Name)
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

	if !alliance.CheckPermission(conn.Name, test_constant.MEMBER_PERMISSION_ADMIN) {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_15}
	}

	return alliance.DestoryItem(info.Index)
}

//整理物品
func ClearUp(conn *test_net.ConnTest, message *test_net.Message) *test_model.ResponseInfo {
	character := test_common.CacheMap.CharacterGetByName(conn.Name)
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

	if !alliance.CheckPermission(conn.Name, test_constant.MEMBER_PERMISSION_ADMIN) {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_15}
	}

	alliance.ClearUp()

	return &test_model.ResponseInfo{Code: test_constant.RES_OK}
}
