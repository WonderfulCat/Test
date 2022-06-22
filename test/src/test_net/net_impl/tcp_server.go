package net_impl

import (
	"fmt"
	"net"
	"test/src/test_net/net_interface"
)

//iServer 接口实现，定义一个Server服务类
type TcpServer struct {
	//tcp4 or other
	IPVersion string
	//服务绑定的IP地址
	IP string
	//服务绑定的端口
	Port int
	//当前Server的消息管理模块，用来绑定MsgId和对应的处理方法
	msgHandler net_interface.MsgHandleI
	//该Server的连接创建时Hook函数
	OnConnStart func(conn net_interface.ConnectionI)
	//该Server的连接断开时的Hook函数
	OnConnStop func(conn net_interface.ConnectionI)
}

func NewServer(ip string, port int) net_interface.ServerI {
	s := &TcpServer{
		IPVersion:  "tcp4",
		IP:         ip,
		Port:       port,
		msgHandler: NewMsgHandle(),
	}
	return s
}

//开启网络服务
func (s *TcpServer) Start() {
	fmt.Printf("[START]  listenner at IP: %s, Port %d is starting\n", s.IP, s.Port)

	//开启一个go去做服务端Linster业务
	go func() {
		//0 启动worker工作池机制
		go s.msgHandler.StartWorker()

		//1 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}

		//2 监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}

		//已经监听成功
		fmt.Println("start server success , now listenning...")

		//3 启动server网络连接业务
		for {
			//3.1 阻塞等待客户端建立连接请求
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}
			fmt.Println("Get conn remote addr = ", conn.RemoteAddr().String())

			//3.3 处理该新连接请求的 业务 方法， 此时应该有 handler 和 conn是绑定的
			dealConn := NewConnTest(s, conn, s.msgHandler)

			//3.4 启动当前链接的处理业务
			go dealConn.Start()
		}
	}()
}

//停止服务
func (s *TcpServer) Stop() {
	fmt.Println("[STOP]  server ")

}

//路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
func (s *TcpServer) AddRouter(msgId int32, router net_interface.RouterI) {
	s.msgHandler.AddRouter(msgId, router)
}

//设置该Server的连接创建时Hook函数
func (s *TcpServer) SetOnConnStart(hookFunc func(net_interface.ConnectionI)) {
	s.OnConnStart = hookFunc
}

//设置该Server的连接断开时的Hook函数
func (s *TcpServer) SetOnConnStop(hookFunc func(net_interface.ConnectionI)) {
	s.OnConnStop = hookFunc
}

//调用连接OnConnStart Hook函数
func (s *TcpServer) CallOnConnStart(conn net_interface.ConnectionI) {
	if s.OnConnStart != nil {
		fmt.Println("---> CallOnConnStart....")
		s.OnConnStart(conn)
	}
}

//调用连接OnConnStop Hook函数
func (s *TcpServer) CallOnConnStop(conn net_interface.ConnectionI) {
	if s.OnConnStop != nil {
		fmt.Println("---> CallOnConnStop....")
		s.OnConnStop(conn)
	}
}
