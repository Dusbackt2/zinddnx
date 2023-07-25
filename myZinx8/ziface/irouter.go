package ziface

/**
路由抽象接口 路由里的数据都是irequest请求
*/

type IRouter interface {
	//处理conn业务之前 hook
	PreHandle(request IRequest)
	//处理业务主
	Handle(request IRequest)
	//处理conn业务之后的hook
	PostHandle(request IRequest)
}
