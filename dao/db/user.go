package db

import (
	"github.com/lihuicms-code-rep/texaspoker/log"
	"github.com/lihuicms-code-rep/texaspoker/model"
	"github.com/pkg/errors"
)

//dao层,user表相关操作

//建表,目前仅仅只有user
func CreateUserTable() error {
	sqlStr := `CREATE TABLE user` +
		`(
                uid INT(10) NOT NULL AUTO_INCREMENT,
                name VARCHAR(255) NOT NULL,
                password VARCHAR(255) NOT NULL,
                create_time DATETIME DEFAULT NULL,
                login_time DATETIME DEFAULT NULL,
                logout_time DATETIME DEFAULT NULL,
                PRIMARY KEY(uid)
        )ENGINE=InnoDB AUTO_INCREMENT=31798340 DEFAULT CHARSET=utf8;`

	result, err := DB.Exec(sqlStr)
	if err != nil {
		log.HttpConsole.Errorf("create table:user error:%+v", err)
		return err
	}

	log.HttpConsole.Infof("create table:user success, result:%+v", result)
	return nil
}

//插入一条玩家数据
func InsertUser(u *model.User) (int64, error) {
	if u == nil {
		return -1, errors.New("user is nil")
	}

	log.HttpConsole.Infof("insert user:%+v", u)

	sqlStr := "insert into user(name, password, create_time, login_time, logout_time)" +
		      "values(?,?,?,?,?)"

	result, err := DB.Exec(sqlStr, u.Name, u.Password, u.CreateTime, u.LoginTime, u.LogoutTime)
	if err != nil {
		log.HttpConsole.Errorf("insert user table error:%+v", err)
		return -1, errors.New("insert user table error")
	}

	uid, err := result.LastInsertId()
	return uid, err
}


//根据玩家id查询玩家信息
//同时也可以作为判断该玩家是否存在的判断
func GetUserByID(uid int64) (*model.User, error) {
	if uid < 0 {
		return nil, errors.New("uid illegal")
	}

	sqlStr := "select uid, user_name, password, create_time, login_time, logout_time from user" +
		"where uid = ?"

	user := model.NewUser()
	err := DB.Get(user, sqlStr, uid)
	if err != nil {
		return nil, errors.New("get user by id error")
	}

	//如果查出来user部分为默认值,作为该id不存在的情况
	if user.UID == 0 {
		return nil, nil
	}

	return user, nil
}
