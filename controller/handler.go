package controller

import (
	"fmt"
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
	log.HttpConsole.Infof("user register......")
	user := model.NewUser() //JSON数据绑定对象
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
//当然还需要修改部分字段:最后一次登录时间
func UserLogin(c *gin.Context) {
	log.HttpConsole.Info("user login......")
	user := model.NewUser() //JSON数据绑定对象
	if err := c.BindJSON(user); err != nil {
		log.HttpConsole.Errorf("bind json data error:%+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "bind json data error",
			"code": defination.ErrorCodeBindJSONFailed,
		})
		return
	}

	success, name := db.CheckLogin(user.UID, user.Password)
	if !success {
		c.JSON(http.StatusOK, gin.H{
			"msg": "login failure",
			"code": defination.ErrorCodeLoginFailed,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg" : fmt.Sprintf("%s login success", name),
		"code" : defination.ErrorCodeSuccess,
	})
}

//用户登出
func UserLogout(c *gin.Context) {

}
