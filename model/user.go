package model

import "time"

//mysql存储Model层
type User struct {
	UID        int64     `db:"uid"`         //用户ID
	UserName   string    `db:"user_name"`   //用户名
	Password   string    `db:"password"`    //密码
	CreateTime time.Time `db:"create_time"` //头像
	LoginTime  time.Time `db:"login_time"`  //最后登录时间
	LogoutTime time.Time `db:"logout_time"` //最后登出时间
}

func NewUser() *User {
	return &User{}
}
