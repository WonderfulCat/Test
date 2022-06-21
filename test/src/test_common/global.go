package test_common

import (
	"reflect"

	"test/src/test_constant"
	"test/src/test_interface"
)

var ClassReflect = map[string]reflect.Type{}

var CacheMap test_interface.CacheI

//注册struct
func RegisterI(name string, classType reflect.Type) {
	ClassReflect[name] = classType
}

//获取struct
func GetReflectByName(name string) interface{} {
	if t, ok := ClassReflect[name]; ok {
		return reflect.New(t).Interface()
	}
	return nil
}

func InitCache() {
	if CacheMap != nil {
		return
	}

	CacheMap = GetReflectByName(test_constant.REGISTER_NAME_CACHE).(test_interface.CacheI)
	CacheMap.Build()
}
