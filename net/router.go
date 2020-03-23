package net

import (
	"github.com/gin-gonic/gin"
	"github.com/lihuicms-code-rep/texaspoker/controller"
	"github.com/lihuicms-code-rep/texaspoker/pb"
	"github.com/lihuicms-code-rep/zinx/ziface"
)

//tcp路由映射

var tcpRouters = map[pb.GameMessage]ziface.IRouter{
	pb.GameMessage_Req_EnterTable: nil,
}


//http路由映射
var getRouters = map[string]func(c *gin.Context){
	"/":         controller.Index,
}

var postRouters = map[string]func(c *gin.Context){
	"/user/register": controller.UserRegister,
	"/user/login":    controller.UserLogin,
	"/user/logout":   controller.UserLogout,
}
