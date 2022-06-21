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

//登陆 (不存在则创建)
func Login(conn *test_net.ConnTest, message *test_net.Message) *test_model.ResponseInfo {
	info := &test_model.LoginRequestInfo{}
	if err := json.Unmarshal(message.GetData(), info); err != nil {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: fmt.Sprintf(test_constant.RES_ERR_MSG_16, err.Error())}
	}

	_, res := test_common.CacheMap.CharacterGetByNamePswd(info.Name, info.Pswd)

	//未注册则注册
	if res.Code == test_constant.RES_REGISTER {
		//创建
		character := test_common.GetReflectByName(test_constant.REGISTER_NAME_CHARACTER).(test_interface.CharacterI)
		character.Build(info.Name, info.Pswd)
		//缓存
		test_common.CacheMap.CharacterAdd(character)
		_, res = test_common.CacheMap.CharacterGetByNamePswd(info.Name, info.Pswd)

	}

	if res.Code == test_constant.RES_OK {
		conn.Name = info.Name
	}

	return res
}
