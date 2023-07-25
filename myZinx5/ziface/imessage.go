package ziface

/*
将请求的消息封装到message中，定义抽象的接口
*/
type IMessage interface {
	GetMsgId() uint32
	GetMsgLen() uint32
	GetMsgData() []byte
	SetMsgId(uint32)
	SetMsgLen(uint32)
	SetData([]byte)
}
