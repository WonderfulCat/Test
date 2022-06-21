package test_impl

import (
	"test/src/test_constant"
	"test/src/test_interface"
	"test/src/test_model"
)

type CacheCharacter struct {
	CharacterMap map[string]test_interface.CharacterI
}

func (c *CacheCharacter) Build() {
	c.CharacterMap = make(map[string]test_interface.CharacterI)

}

func (c *CacheCharacter) CharacterGetByNamePswd(name, pswd string) (test_interface.CharacterI, *test_model.ResponseInfo) {
	if ret, ok := c.CharacterMap[name]; ok {
		if ret.GetPswd() == pswd {
			return ret, &test_model.ResponseInfo{Code: test_constant.RES_OK}
		}
		return nil, &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_8}
	}

	return nil, &test_model.ResponseInfo{Code: test_constant.RES_REGISTER}
}

func (c *CacheCharacter) CharacterGetByName(name string) test_interface.CharacterI {
	if ret, ok := c.CharacterMap[name]; ok {
		return ret
	}

	return nil
}

func (c *CacheCharacter) CharacterAdd(character test_interface.CharacterI) bool {
	c.CharacterMap[character.GetName()] = character
	return true
}
