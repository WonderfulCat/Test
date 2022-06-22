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

type LoginHandle struct {
	net_impl.BaseRouter
}

func (c *LoginHandle) Handle(req net_interface.RequestI) {
	ret := c.Login(req)
	if err := req.GetConnection().SendBuffMsg(req.GetMsgID(), GetJsonBytes(ret)); err != nil {
		fmt.Println(err)
	}
}

//登陆 (不存在则创建)
func (c *LoginHandle) Login(request net_interface.RequestI) *test_model.ResponseInfo {
	info := &test_model.LoginRequestInfo{}
	if err := json.Unmarshal(request.GetData(), info); err != nil {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: fmt.Sprintf(test_constant.RES_ERR_MSG_16, err.Error())}
	}

	_, code := test_common.CacheMap.CharacterGetByNamePswd(info.Name, info.Pswd)

	switch code {
	case test_constant.RES_REGISTER: //未注册则注册
		//创建
		character := GetReflectByName(test_constant.REGISTER_NAME_CHARACTER).(test_interface.CharacterI)
		character.Build(info.Name, info.Pswd)
		//缓存
		if !test_common.CacheMap.CharacterAdd(character) {
			return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_25}
		}
	case test_constant.RES_ERR:
		request.GetConnection().SetName("")
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_8}

	}

	request.GetConnection().SetName(info.Name)

	return &test_model.ResponseInfo{Code: test_constant.RES_OK, Msg: fmt.Sprintf(test_constant.RES_ERR_MSG_18, info.Name)}

}
