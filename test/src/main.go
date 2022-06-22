package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"test/src/test_net/net_impl"
	"test/src/test_service"
)

func main() {
	//初始化服务器
	s := net_impl.NewServer("0.0.0.0", 8080)
	//初始化struct
	test_service.InitStructRouter()
	//初始化缓存
	test_service.InitCache()
	//初始化router
	test_service.InitRouter(s)
	//初始化data
	test_service.InitItemData("testItem.data")
	//启动
	s.Start()

	//关闭HOOKS
	WaitExit(closeServer)
}

func WaitExit(beforeExit ...func()) {
	catch := make(chan os.Signal, 1)

	signal.Notify(catch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	for sig := range catch {
		fmt.Printf("System signal is caught (%v) \n", sig.String())
		break
	}

	close(catch)

	for _, fn := range beforeExit {
		fn()
	}
}

//close
func closeServer() {
	//关闭前保存数据
	fmt.Println("Server Stopped")
}
