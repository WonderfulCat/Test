package test_net

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"time"
)

func HandleClient(conn net.Conn) {

	//声明一个管道用于接收解包的数据
	readChannel := make(chan *Message, 16)
	go Read(NewConnTest(conn, ""), readChannel, 64)

	for {
		// 创建拆包解包的对象
		dp := NewDataPack()

		//读取客户端的Msg head
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, headData); err != nil {
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
			if _, err := io.ReadFull(conn, data); err != nil {
				fmt.Println("read msg data error ", err)
				break
			}
		}
		msg.SetData(data)

		readChannel <- msg
	}
}

func Read(conn *ConnTest, readeChannel <-chan *Message, timeout int) {
	for {
		select {
		case data := <-readeChannel:
			conn.Conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
			Deal(conn, data)
			break
		case <-time.After(time.Duration(timeout) * time.Second):
			conn.Conn.Close()
			fmt.Println("connection is closed.")
			return
		}
	}
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
*/
func Deal(conn *ConnTest, message *Message) {
	f, ok := Routers[message.GetMsgId()]
	if !ok {
		return
	}

	//处理
	ret := f(conn, message)

	//json
	data, err := json.Marshal(ret)
	if err != nil {
		fmt.Println(err)
		return
	}

	//将data封包，并且发送
	dp := NewDataPack()
	msg, err := dp.Pack(NewMsgPackage(message.GetMsgId(), data))
	if err != nil {
		fmt.Println(err)
		return
	}

	conn.Conn.Write(msg)
}
