package net

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lihuicms-code-rep/texaspoker/config"
	"github.com/lihuicms-code-rep/texaspoker/log"
)

//处理用户注册,登入,登出及后续需要http处理的内容

func StartHttpServer() {
	switch config.HTTPServerConfig.Mode {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()
	for path, handler := range getRouters {
		r.GET(path, handler)
	}

	for path, handler := range postRouters {
		r.POST(path, handler)
	}

	addr := fmt.Sprintf("%s:%d", config.HTTPServerConfig.Host, config.HTTPServerConfig.Port)
	if err := r.Run(addr); err != nil {
		log.Console.Errorf("http server run and serve on %s error:%+v", addr, err)
		return
	}

	log.HttpConsole.Infof("http server run and server on %s", addr)
}
