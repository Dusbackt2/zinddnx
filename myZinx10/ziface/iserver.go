package ziface

type IServer interface {
	Start()

	Serve()

	Stop()

	AddRouter(msgId uint32, router IRouter)

	GetConnManager() IConnManager

	SetOnConnStart(func(connection IConnection))

	SetOnConnStop(func(connection IConnection))

	CallOnConnStart(connection IConnection)

	CallOnConnStop(connection IConnection)
}
