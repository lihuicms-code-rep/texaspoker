package config

//一个类型桌子基本配置
type TableOneType struct {
	Type        int32   `json:"type"`          //桌子类型
	Level       int32   `json:"level"`         //桌子盲注级别
	MaxGamer    uint8   `json:"max_gamer"`     //最大游戏人数
	MaxWatcher  uint8   `json:"max_watcher"`   //最大围观人数
	Num         int32   `json:"num"`           //桌子数量
	StartID     int32   `json:"start_id"`      //起始桌子id
	UseRobotTID []int32 `json:"use_robot_tid"` //允许投入机器人的桌子编号
}

//多种类型桌子组成一个房间
type Room struct {
	TableList []TableOneType `json:"table_list"`
}

