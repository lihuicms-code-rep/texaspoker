package model

import "time"

//mysql存储Model层
//tag标签,db:sqlx模块中对应数据库表字段名,form,binding:gin中使用的
type User struct {
	UID        uint32    `db:"uid" json:"uid"`                           //用户ID
	Name       string    `db:"name" json:"name" form:"name"`             //用户名
	Password   string    `db:"password" json:"password" form:"password"` //密码
	CreateTime time.Time `db:"create_time" json:"create_name"`           //创建时间
	LoginTime  time.Time `db:"login_time" json:"login_name"`             //最后登录时间
	LogoutTime time.Time `db:"logout_time" json:"logout_name"`           //最后登出时间
}

func NewUser() *User {
	return &User{
		UID:        0,
		Name:       "",
		Password:   "",
		CreateTime: time.Now(),
		LoginTime:  time.Now(),
		LogoutTime: time.Now(),
	}
}
