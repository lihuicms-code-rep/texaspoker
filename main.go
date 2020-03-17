package main

import (
	"flag"
	"github.com/lihuicms-code-rep/texaspoker/config"
	"github.com/lihuicms-code-rep/texaspoker/dao/db"
	"github.com/lihuicms-code-rep/texaspoker/dao/redis"
	"github.com/lihuicms-code-rep/texaspoker/log"
	"github.com/lihuicms-code-rep/texaspoker/net"
)

//启动参数
var (
	configPath  string //配置路径
)

func init() {
	//命令行参数
	flag.StringVar(&configPath, "config", "./res/", "game config path")
	flag.Parse()

	//logger初始化
	log.InitZapLogger()

	//config Load
	config.Load(configPath)

	//mysql连接
	db.Connect(config.GetMySQLDSN())

	//redis客户端初始化
	redis.InitRedisClient(config.DBConfig.Redis.Address, config.DBConfig.Redis.Password, config.DBConfig.Redis.DBNumber)
}

func main() {
	//网络启动
	net.StartServer()
}
