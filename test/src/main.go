package main

import (
	"fmt"
	"net"
	"reflect"
	"test/src/test_common"
	"test/src/test_constant"
	"test/src/test_impl"
	"test/src/test_net"
	"test/src/test_service"
)

func init() {
	test_common.RegisterI(test_constant.REGISTER_NAME_CACHE, reflect.TypeOf((*test_impl.Cache)(nil)).Elem())
	test_common.RegisterI(test_constant.REGISTER_NAME_CHARACTER, reflect.TypeOf((*test_impl.CharacterInfo)(nil)).Elem())
	test_common.RegisterI(test_constant.REGISTER_NAME_ALLIANCE, reflect.TypeOf((*test_impl.AllianceInfo)(nil)).Elem())

	//router
	test_net.RegisterRouter(1000, test_service.Login)
	test_net.RegisterRouter(1001, test_service.WhichAlliance)
	test_net.RegisterRouter(1002, test_service.CreateAlliance)
	test_net.RegisterRouter(1003, test_service.JoinAlliance)
	test_net.RegisterRouter(1004, test_service.DismissAlliance)
	test_net.RegisterRouter(1005, test_service.IncreaseCapacity)
	test_net.RegisterRouter(1006, test_service.StoreItem)
	test_net.RegisterRouter(1007, test_service.DestoryItem)
	test_net.RegisterRouter(1008, test_service.ClearUp)
	test_net.RegisterRouter(1009, test_service.AllianceList)
}

func main() {
	test_service.InitCache()

	listener, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		fmt.Println("服务器启动: ", "0.0.0.0:8081")

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("监听错误: ", err.Error())
			continue
		}

		go test_net.HandleClient(conn)
	}

}
