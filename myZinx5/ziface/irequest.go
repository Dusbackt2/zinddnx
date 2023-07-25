package ziface

/**
实际上 客户端请求的连接信息和请求数据 包装
*/

type IRequest interface {
	//得到当前连接
	GetConnection() IConnection
	//得到请求的消息数据
	GetData() []byte

	GetMessageId() uint32
}
