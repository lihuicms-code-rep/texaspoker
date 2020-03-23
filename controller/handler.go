package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lihuicms-code-rep/texaspoker/dao/db"
	"github.com/lihuicms-code-rep/texaspoker/defination"
	"github.com/lihuicms-code-rep/texaspoker/log"
	"github.com/lihuicms-code-rep/texaspoker/model"
	"net/http"
)

//HTTP路由处理层

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"data": "welcome to texaspoker web",
	})
}

//用户注册
func UserRegister(c *gin.Context) {
	user := model.NewUser() //JSON数据绑定对象
	log.HttpConsole.Infof("init user:%+v", user)
	if err := c.BindJSON(user); err != nil {
		log.HttpConsole.Errorf("bind json data error:%+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "bind json data error",
			"code": defination.ErrorCodeBindJSONFailed,
		})
		return
	}

	uid, err := db.InsertUser(user)
	if err != nil {
		log.HttpConsole.Errorf("insert user error:%+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "register user failed",
			"code": defination.ErrorCodeInsertUserFailed,
		})
		return
	}

	log.HttpConsole.Infof("uid:%d register success", uid)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"code": defination.ErrorCodeSuccess,
		"uid":  uid,
	})
}

//用户登录
func UserLogin(c *gin.Context) {


}

//用户登出
func UserLogout(c *gin.Context) {

}
