package net

import (
	"github.com/lihuicms-code-rep/texaspoker/config"
	"github.com/lihuicms-code-rep/texaspoker/log"
	"github.com/lihuicms-code-rep/zinx/ziface"
	"github.com/lihuicms-code-rep/zinx/znet"
)

func StartTCPServer() {
	var server ziface.IServer
	server = znet.NewServer()
	log.Console.Infof("server:%s start and serve", config.TCPServerConfig.Name)

	//注册Hook函数
	server.SetOnConnStart(OnConnectionStart)
	server.SetOnConnStop(OnConnectionLost)

	//注册业务路由
	for id, router := range tcpRouters {
		server.AddRouter(uint32(id), router)
	}

	//启动服务
	server.Serve()
}