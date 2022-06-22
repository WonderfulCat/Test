package test_service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"test/src/test_common"
	"test/src/test_constant"
	"test/src/test_impl"
	"test/src/test_interface"
	"test/src/test_net/net_interface"
	"test/src/test_pb"

	"google.golang.org/protobuf/proto"
)

func InitStructRouter() {
	test_common.ClassReflect = make(map[string]reflect.Type)

	RegisterI(test_constant.REGISTER_NAME_CACHE, reflect.TypeOf((*test_impl.Cache)(nil)).Elem())
	RegisterI(test_constant.REGISTER_NAME_CHARACTER, reflect.TypeOf((*test_impl.CharacterInfo)(nil)).Elem())
	RegisterI(test_constant.REGISTER_NAME_ALLIANCE, reflect.TypeOf((*test_impl.AllianceInfo)(nil)).Elem())
}

//注册struct
func RegisterI(name string, classType reflect.Type) {
	test_common.ClassReflect[name] = classType
}

//获取struct
func GetReflectByName(name string) interface{} {
	if t, ok := test_common.ClassReflect[name]; ok {
		return reflect.New(t).Interface()
	}
	return nil
}

func InitCache() {
	if test_common.CacheMap != nil {
		return
	}

	test_common.CacheMap = GetReflectByName(test_constant.REGISTER_NAME_CACHE).(test_interface.CacheI)
	test_common.CacheMap.Build()
}

func InitItemData(path string) {
	dir, _ := os.Getwd()

	in, err := ioutil.ReadFile(dir + "/" + path)
	if err != nil {
		log.Fatalln("Error reading file:", err)

	}

	items := &test_pb.TestItem_Array{}

	if err := proto.Unmarshal(in, items); err != nil {
		log.Fatalln("Failed to parse itemdata:", err)

	}

	test_common.ItemData = items
}

/**
op_code
1000 : login
1001 : whichAlliance
1002 : createAlliance
1003 : joinAlliance
1004 : dismissAlliance
1005 : increaseCapacity
1006 : storeItem
1007 : destoryItem
1008 : clearUp
1009 : allianceList

*/
func InitRouter(server net_interface.ServerI) {
	server.AddRouter(1000, &LoginHandle{})
	server.AddRouter(1001, &WhichAllianceHandle{})
	server.AddRouter(1002, &CreateAllianceHandle{})
	server.AddRouter(1003, &JoinAllianceHandle{})
	server.AddRouter(1004, &DismissAllianceHandle{})
	server.AddRouter(1005, &IncreaseCapacityHandle{})
	server.AddRouter(1006, &StoreItemHandle{})
	server.AddRouter(1007, &DestoryItemHandle{})
	server.AddRouter(1008, &ClearUpHandle{})
	server.AddRouter(1009, &AllianceListHandle{})
	server.AddRouter(1010, &GetItemListHandle{})

	server.SetOnConnStart(OnConnection)
	server.SetOnConnStop(OnDisConnection)
}

//连接建立时执行
func OnConnection(conn net_interface.ConnectionI) {
	fmt.Println(conn.GetTCPConnection().RemoteAddr().String() + " is connection.")
}

//连接断开的时候执行
func OnDisConnection(conn net_interface.ConnectionI) {
	fmt.Println(conn.GetTCPConnection().RemoteAddr().String() + " is disconnection.")
}

func GetJsonBytes(message interface{}) []byte {
	//json
	data, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return data
}
