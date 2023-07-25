package ziface

type IConnManager interface {
	//添加连接
	Add(connection IConnection)
	//删除连接
	Remove(connection IConnection)
	//根据connId获取连接
	Get(connId uint32) (IConnection, error)
	//得到总数
	Len() int
	//清除所有连接
	ClearConn()
}
