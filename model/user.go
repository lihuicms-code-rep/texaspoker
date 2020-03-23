package model

import "time"

//mysql存储Model层
type User struct {
	UID        int64     `db:"uid" json:"uid"`                                              //用户ID
	Name       string    `db:"name" json:"name" form:"name" binding:"required"`             //用户名
	Password   string    `db:"password" json:"password" form:"password" binding:"required"` //密码
	CreateTime time.Time `db:"create_time" json:"user_name"`                                //创建时间
	LoginTime  time.Time `db:"login_time" json:"user_name"`                                 //最后登录时间
	LogoutTime time.Time `db:"logout_time" json:"user_name"`                                //最后登出时间
}

func NewUser() *User {
	return &User{
		UID:        -1,
		Name:       "",
		Password:   "",
		CreateTime: time.Now(),
		LoginTime:  time.Now(),
		LogoutTime: time.Now(),
	}
}
