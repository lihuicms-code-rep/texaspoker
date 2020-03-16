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
)

//配置名,配置示例对应关系
var configNames = map[string]interface{}{
	"server" : ServerConfig,
}

//Load配置
func Load(path string) {
	for name, cfgObj := range configNames {
		_, err := utils.ReadFile2Obj(getFileName(path, name), cfgObj)
		if err != nil {
			break
		}
	}
	log.Console.Info("Load Config Ok")
}

func getFileName(path, name string) string {
	return fmt.Sprintf("%s%s.json", path, name)
}



