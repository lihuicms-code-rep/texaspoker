package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//HTTP路由处理层

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg" : "ok",
		"data" : "welcome to texaspoker web",
	})
}


//用户注册
func UserRegister(c *gin.Context) {

}

//用户登录
func UserLogin(c *gin.Context) {

}

//用户登出
func UserLogout(c *gin.Context) {

}