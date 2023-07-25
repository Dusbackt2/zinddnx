package znet

import (
	"fmt"
	"myZinx7/ziface"
)

type MsgHandle struct {
	Apis map[uint32]ziface.IRouter
}

func NewMsgHandle() *MsgHandle {
	m := &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
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
