package znet

import (
	"fmt"
	"myZinx7/ziface"
)

// 初衷实现router时先嵌入这个基类，根据需求重写
type BaseRouter struct{}

func (br *BaseRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("preHandle")
}

func (br *BaseRouter) Handle(request ziface.IRequest) {
	fmt.Println("handle")
}

func (br *BaseRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("postHandle")
}
