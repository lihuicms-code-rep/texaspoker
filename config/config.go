package config

import (
	"fmt"
	"github.com/lihuicms-code-rep/texaspoker/log"
	"github.com/lihuicms-code-rep/texaspoker/utils"
)

//游戏配置读取

//配置对象
var (
	TCPServerConfig  = NewTCPServerConfig()
	HTTPServerConfig = NewHTTPServerConfig()
	DBConfig         = NewDBConfig()
)

//配置名,配置实例对应关系
var configNames = map[string]interface{}{
	"tcpserver":  TCPServerConfig,
	"httpserver": HTTPServerConfig,
	"db":         DBConfig,
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
	log.Console.Infof("Load Config Ok, TCPServerConfig:%+v", TCPServerConfig)
	log.Console.Infof("Load Config Ok, HTTPServerConfig:%+v", HTTPServerConfig)
	log.Console.Infof("Load Config Ok, DBConfig:%+v", DBConfig)
}

func getFileName(path, name string) string {
	return fmt.Sprintf("%s%s.json", path, name)
}

//获取MySQL连接串
//dsn串结构:[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
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
