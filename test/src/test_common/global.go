package test_common

import (
	"reflect"

	"test/src/test_interface"
	"test/src/test_pb"
)

var (
	ClassReflect map[string]reflect.Type
	CacheMap     test_interface.CacheI
	ItemData     *test_pb.TestItem_Array
)
