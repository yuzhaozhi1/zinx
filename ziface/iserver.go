// Package ziface 主要提供zinx全部抽象层接口定义.
// 包括:
//		IServer 服务mod接口
//		IRouter 路由mod接口
//		IConnection 连接mod层接口
//      IMessage 消息mod接口
//		IDataPack 消息拆解接口
//      IMsgHandler 消息处理及协程池接口

package ziface

type IServer interface {
	Start() // 启动服务器的方法
	Stop()  // 停止服务器的方法
	Serve() // 开启业务服务的方法
}
