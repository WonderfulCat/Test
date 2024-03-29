package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"test/src/test_model"
	"test/src/test_net/net_impl"
)

var address = flag.String("ip", "0.0.0.0:8080", "请按格式输入: IP:PORT")

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", *address)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		return
	}

	var msg string

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("=============指令列表如下,空格分割参数============= ")
		fmt.Println("/login name pswd")
		fmt.Println("/whichAlliance")
		fmt.Println("/createAlliance allianceName")
		fmt.Println("/allianceList")
		fmt.Println("/joinAlliance name")
		fmt.Println("/dismissAlliance")
		fmt.Println("/increaseCapacity")
		fmt.Println("/storeItem itemId itemNum index")
		fmt.Println("/destroyItem index")
		fmt.Println("/clearup")
		fmt.Println("/getItemList")
		fmt.Println("请输入需要操作的指令: ")
		msg, _ = reader.ReadString('\n')

		switch SendMsgData(conn, msg) {
		case -1:
			fmt.Println("指令输入错误.")
			continue
		case -2:
			fmt.Println("数据错误.")
			continue
		case -3:
			fmt.Println("连接错误,退出客户端.")
			return

		}

		message := HandleClient(conn)
		fmt.Println("---------------返回结果-------------")
		if message == nil {
			fmt.Println("返回数据为空")
		} else {
			fmt.Println(message.Id, string(message.Data))
		}

		fmt.Println()
		fmt.Println()
	}
}

func SendMsgData(conn net.Conn, msg string) int {
	if len(msg) <= 0 {
		return -1
	}

	var msgId int32
	var data []byte

	ss := strings.Split(msg, " ")
	switch strings.TrimSpace(ss[0]) {
	case "/login":
		if len(ss) != 3 {
			return -1
		}

		msgId = 1000
		data, _ = json.Marshal(&test_model.LoginRequestInfo{Name: strings.TrimSpace(ss[1]), Pswd: strings.TrimSpace(ss[2])})
	case "/whichAlliance":
		msgId = 1001
		data, _ = json.Marshal(&test_model.WhichAllianceRequestInfo{})
	case "/createAlliance":
		if len(ss) != 2 {
			return -1
		}

		msgId = 1002
		data, _ = json.Marshal(&test_model.CreateAllianceRequestInfo{AName: strings.TrimSpace(ss[1])})
	case "/joinAlliance":
		if len(ss) != 2 {
			return -1
		}

		msgId = 1003
		data, _ = json.Marshal(&test_model.JoinAllianceRequestInfo{AName: strings.TrimSpace(ss[1])})
	case "/dismissAlliance":
		msgId = 1004
		data, _ = json.Marshal(&test_model.DismissAllianceRequestInfo{})
	case "/increaseCapacity":
		msgId = 1005
		data, _ = json.Marshal(&test_model.DismissAllianceRequestInfo{})
	case "/storeItem":
		if len(ss) != 4 {
			return -1
		}

		msgId = 1006

		itemId, _ := strconv.ParseInt(strings.TrimSpace(ss[1]), 10, 64)
		itemNum, _ := strconv.ParseInt(strings.TrimSpace(ss[2]), 10, 64)
		itemIndex, _ := strconv.ParseInt(strings.TrimSpace(ss[3]), 10, 64)
		data, _ = json.Marshal(&test_model.StoreItemRequestInfo{ItemId: int32(itemId), ItemNum: int32(itemNum), Index: int32(itemIndex)})
	case "/destroyItem":
		if len(ss) != 2 {
			return -1
		}

		itemIndex, _ := strconv.ParseInt(strings.TrimSpace(ss[1]), 10, 64)
		msgId = 1007
		data, _ = json.Marshal(&test_model.DestoryItemRequestInfo{Index: int32(itemIndex)})
	case "/clearup":
		msgId = 1008
		data, _ = json.Marshal(&test_model.ClearUpRequestInfo{})
	case "/allianceList":
		msgId = 1009
		data, _ = json.Marshal(&test_model.AllianceList{})
	case "/getItemList":
		msgId = 1010
		data, _ = json.Marshal(&test_model.GetItemList{})
	default:
		fmt.Println("无效的指令")
		return -1
	}

	dp := net_impl.NewDataPack()
	sendMsg, err := dp.Pack(net_impl.NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println(err)
		return -2
	}

	_, err = conn.Write(sendMsg)
	if err != nil {
		fmt.Println(err)
		return -3
	}

	return 0
}

func HandleClient(conn net.Conn) *net_impl.Message {
	// 创建拆包解包的对象
	dp := net_impl.NewDataPack()

	//读取客户端的Msg head
	headData := make([]byte, dp.GetHeadLen())
	if _, err := io.ReadFull(conn, headData); err != nil {
		fmt.Println("read msg head error ", err)
		return nil
	}

	//拆包，得到msgid 和 datalen 放在msg中
	msg, err := dp.Unpack(headData)
	if err != nil {
		fmt.Println("unpack error ", err)
		return nil
	}

	//根据 dataLen 读取 data，放在msg.Data中
	//长度包含id  需要减去已经取得的id的长度4字节
	var data []byte
	if msg.GetDataLen() > 0 {
		data = make([]byte, msg.GetDataLen()-4)
		if _, err := io.ReadFull(conn, data); err != nil {
			fmt.Println("read msg data error ", err)
			return nil
		}
	}
	msg.SetData(data)

	return msg
}
