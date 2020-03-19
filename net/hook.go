package net

import (
	"github.com/lihuicms-code-rep/texaspoker/log"
	"github.com/lihuicms-code-rep/zinx/ziface"
)
//TCP连接建立后的HOOK函数

//连接建立后要做的工作
func OnConnectionStart(conn ziface.IConnection) {
	log.Console.Info("register connection start func")
}


//连接断开前要做的工作
func OnConnectionLost(conn ziface.IConnection) {
	log.Console.Info("register connection lost func")
}
