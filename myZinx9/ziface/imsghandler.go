package ziface

/**
消息管理
*/

type IMsgHandle interface {
	//调度/执行对应的Router消息处理方法
	DoMsgHandler(request IRequest)
	//为消息添加具体的处理逻辑
	AddRouter(msgId uint32, router IRouter)

	StartWorkerPool()

	SendMsgToTaskQueue(request IRequest)
}