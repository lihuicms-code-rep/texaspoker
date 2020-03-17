package net

import (
	"github.com/lihuicms-code-rep/texaspoker/pb"
	"github.com/lihuicms-code-rep/zinx/ziface"
)

//路由处理

var routers = map[pb.GameMessage]ziface.IRouter{
	pb.GameMessage_Req_Login:      nil,
	pb.GameMessage_Req_EnterTable: nil,
}
