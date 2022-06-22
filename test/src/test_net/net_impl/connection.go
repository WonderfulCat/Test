package net_impl

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"test/src/test_net/net_interface"
)

type ConnTest struct {
	TcpServer   net_interface.ServerI
	Conn        *net.TCPConn //原始连接
	Name        string       //自定义数据
	MsgHandler  net_interface.MsgHandleI
	ctx         context.Context
	cancel      context.CancelFunc
	MsgBuffChan chan []byte //写缓存
}

const CHANLEN = 1024

func NewConnTest(tcpServer net_interface.ServerI, conn *net.TCPConn, msgHandler net_interface.MsgHandleI) *ConnTest {
	return &ConnTest{TcpServer: tcpServer, Conn: conn, MsgHandler: msgHandler, MsgBuffChan: make(chan []byte, CHANLEN)}
}

func (c *ConnTest) SetName(name string) {
	c.Name = name
}

func (c *ConnTest) GetName() string {
	return c.Name
}

func (c *ConnTest) StartWriter() {
	for {
		select {
		case data, ok := <-c.MsgBuffChan:
			if ok {
				//有数据要写给客户端
				if _, err := c.Conn.Write(data); err != nil {
					fmt.Println("Send Buff Data error:, ", err, " Conn Writer exit")
					return
				}
			} else {
				fmt.Println("msgBuffChan is Closed")
				return
			}
		case <-c.ctx.Done():
			return
		}
	}
}

func (c *ConnTest) StartReader() {
	defer c.Stop()

	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			// 创建拆包解包的对象
			dp := NewDataPack()

			//读取客户端的Msg head
			headData := make([]byte, dp.GetHeadLen())
			if _, err := io.ReadFull(c.Conn, headData); err != nil {
				fmt.Println("read msg head error ", err)
				return
			}

			//拆包，得到msgid 和 datalen 放在msg中
			msg, err := dp.Unpack(headData)
			if err != nil {
				fmt.Println("unpack error ", err)
				return
			}

			//根据 dataLen 读取 data，放在msg.Data中
			//长度包含id  需要减去已经取得的id的长度4字节
			var data []byte
			if msg.GetDataLen() > 0 {
				data = make([]byte, msg.GetDataLen()-4)
				if _, err := io.ReadFull(c.Conn, data); err != nil {
					fmt.Println("read msg data error ", err)
					break
				}
			}
			msg.SetData(data)

			go c.MsgHandler.SendMsgToTaskQueue(NewRequest(c, msg))
		}
	}
}

//写入消息到buff chan
func (c *ConnTest) SendBuffMsg(msgId int32, data []byte) error {
	//将data封包，并且发送
	dp := NewDataPack()
	msg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println(err)
		return errors.New("Pack error msg ")
	}

	//写回客户端chan
	c.MsgBuffChan <- msg

	return nil
}

//启动连接，让当前连接开始工作
func (c *ConnTest) Start() {
	c.ctx, c.cancel = context.WithCancel(context.Background())
	//1 开启用户从客户端读取数据流程的Goroutine
	go c.StartReader()
	//2 开启用于写回客户端数据流程的Goroutine
	go c.StartWriter()
	//按照用户传递进来的创建连接时需要处理的业务，执行钩子方法
	c.TcpServer.CallOnConnStart(c)
}

func (c *ConnTest) Stop() {
	fmt.Println("Conn Stop()...ConnID = ", c.Name)
	//如果用户注册了该链接的关闭回调业务，那么在此刻应该显示调用
	c.TcpServer.CallOnConnStop(c)
	// 关闭socket链接
	c.Conn.Close()
	// 关闭reader/writer
	c.cancel()
	//关闭该链接全部管道
	close(c.MsgBuffChan)
}

//从当前连接获取原始的socket TCPConn
func (c *ConnTest) GetTCPConnection() *net.TCPConn {
	return c.Conn
}
