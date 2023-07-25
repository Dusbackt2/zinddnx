package znet

import (
	"fmt"
	"myZinx9/ziface"
)

type MsgHandle struct {
	Apis map[uint32]ziface.IRouter
	//负责存放worker取任务的消息队列
	TaskQueue []chan ziface.IRequest
	//业务工作worker池的worker数量
	WorkerPoolSize uint32
}

func NewMsgHandle() *MsgHandle {
	m := &MsgHandle{
		Apis:           make(map[uint32]ziface.IRouter),
		WorkerPoolSize: 8,
		TaskQueue:      make([]chan ziface.IRequest, 8),
	}
	return m
}

func (mh *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	router, ok := mh.Apis[request.GetMessageId()]
	if !ok {
		fmt.Println("api msgId =", request.GetMessageId(), ", is NOT FOUND!need register")
	}
	router.PreHandle(request)
	router.Handle(request)
	router.PostHandle(request)
}

func (mh *MsgHandle) AddRouter(msgId uint32, router ziface.IRouter) {
	if _, ok := mh.Apis[msgId]; ok {
		return
	}
	mh.Apis[msgId] = router
}

//启动一个worker工作池子

func (mh *MsgHandle) StartWorkerPool() {
	fmt.Println("[===start pool===]")
	//根据workerPoolSize分别开启Worker,每个Worker用一个go承载
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		mh.TaskQueue[i] = make(chan ziface.IRequest, 1024)
		go mh.StartOneWorker(uint32(i), mh.TaskQueue[i])
	}

}

// 启动一个worker工作流程
func (mh *MsgHandle) StartOneWorker(workId uint32, taskQueue chan ziface.IRequest) {
	fmt.Println("WorderID=", workId, " is started.....")
	for {
		select {
		case request := <-taskQueue:
			mh.DoMsgHandler(request)
		}
	}
}

func (mh *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {
	fmt.Println("[===send msg===]")
	//将消息平均分配
	u := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	requests := mh.TaskQueue[u]
	//将消息发送给对应的worker的TaskQueue
	requests <- request

}
