package db

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/lihuicms-code-rep/texaspoker/log"
)

//dao层,mysql连接

var (
	DB *sqlx.DB
)

//初始化
func Connect(dsn string) error {
	log.Console.Infof("mysql connect, dsn:%s", dsn)
	database, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Console.Infof("open mysql error:%+v", err)
		return errors.New("open mysql error")
	}
	DB = database

	//测试连接成功与否
	err = DB.Ping()
	if err != nil {
		log.Console.Errorf("ping mysql error:%+v", err)
		return errors.New("ping error")
	}

	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	log.Console.Info("connect to mysql success")
	return nil
}



