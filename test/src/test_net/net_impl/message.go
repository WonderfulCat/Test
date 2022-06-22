package net_impl

type Message struct {
	DataLen uint32 //消息的长度
	Id      int32  //消息的ID
	Data    []byte //消息的内容
}

//创建一个Message消息包
func NewMsgPackage(id int32, data []byte) *Message {
	return &Message{
		DataLen: uint32(len(data) + 4),
		Id:      id,
		Data:    data,
	}
}

//获取消息数据段长度
func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

//获取消息ID
func (msg *Message) GetMsgId() int32 {
	return msg.Id
}

//获取消息内容
func (msg *Message) GetData() []byte {
	return msg.Data
}

//设置消息数据段长度
func (msg *Message) SetDataLen(len uint32) {
	msg.DataLen = len
}

//设计消息ID
func (msg *Message) SetMsgId(msgId int32) {
	msg.Id = msgId
}

//设计消息内容
func (msg *Message) SetData(data []byte) {
	//msgIdData := make([]byte, 4)
	//头四个字节 为消息id
	//再读取msgID(头四个字节)
	//dataBuff := bytes.NewReader(msgIdData)
	//if err := binary.Read(dataBuff, binary.BigEndian, msg.Id); err != nil {
	//	fmt.Println("read msgId error ", err)
	//	return
	//}
	msg.Data = data
}
