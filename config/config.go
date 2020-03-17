package config

import (
	"fmt"
	"github.com/lihuicms-code-rep/texaspoker/log"
	"github.com/lihuicms-code-rep/texaspoker/utils"
)

//游戏配置读取

//配置对象
var (
	ServerConfig = NewServerConfig()
	DBConfig     = NewDBConfig()
)

//配置名,配置示例对应关系
var configNames = map[string]interface{}{
	"server": ServerConfig,
	"db":     DBConfig,
}

//Load配置
func Load(path string) {
	for name, cfgObj := range configNames {
		_, err := utils.ReadFile2Obj(getFileName(path, name), cfgObj)
		if err != nil {
			log.Console.Error("read config:%s error:%v", name, err)
			break
		}
	}
	log.Console.Infof("Load Config Ok, ServerConfig:%+v", ServerConfig)
	log.Console.Infof("Load Config Ok, DBConfig:%+v", DBConfig)
}

func getFileName(path, name string) string {
	return fmt.Sprintf("%s%s.json", path, name)
}

//获取MySQL连接串
func GetMySQLDSN() string {
	dsn := ""
	paramStr := ""
	for k, v := range DBConfig.MySQL.Params {
		s := k + "=" + v + "&"
		paramStr += s
	}

	if paramStr != "" {
		dsn = fmt.Sprintf("%s:%s@%s(%s)/%s?%s",
			DBConfig.MySQL.UserName, DBConfig.MySQL.Password, DBConfig.MySQL.Protocol,
			DBConfig.MySQL.Address, DBConfig.MySQL.DBName, paramStr)
	} else {
		dsn = fmt.Sprintf("%s:%s@%s(%s)/%s",
			DBConfig.MySQL.UserName, DBConfig.MySQL.Password, DBConfig.MySQL.Protocol,
			DBConfig.MySQL.Address, DBConfig.MySQL.DBName)
	}

	return dsn

}
