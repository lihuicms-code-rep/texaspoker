package main

import (
	"flag"
	"github.com/lihuicms-code-rep/texaspoker/config"
	"github.com/lihuicms-code-rep/texaspoker/log"
	"github.com/lihuicms-code-rep/texaspoker/net"
	"os"
	"path/filepath"
)

//启动参数
var (
	configPath  string //配置路径
)

func init() {
	//命令行参数
	flag.StringVar(&configPath, "config", "", "game config path")
	flag.Parse()

	//logger初始化
	log.InitZapLogger()

	//config Load
	config.Load(configPath)
}

//方便日志创建
func getExecPath() string {
	dir, _ := os.Executable()
	return filepath.Dir(dir)
}


func main() {
	//网络启动
	net.StartServer()
}
