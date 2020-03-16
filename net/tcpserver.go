package net

import (
	"github.com/lihuicms-code-rep/texaspoker/config"
	"github.com/lihuicms-code-rep/texaspoker/log"
	"github.com/lihuicms-code-rep/zinx/ziface"
	"github.com/lihuicms-code-rep/zinx/znet"
)

func StartServer() {
	var server ziface.IServer
	server = znet.NewServer()
	server.Serve()
	log.Console.Info("server:%s start and server", config.ServerConfig.Name)
}