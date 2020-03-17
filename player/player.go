package player

import "github.com/lihuicms-code-rep/zinx/ziface"

//玩家数据
type Player struct {
	BaseInfo
	PropInfo
	ConnInfo
}

//基础信息,对应user表
type BaseInfo struct {
}

//道具信息
type PropInfo struct {
	Props []SingleProp
}

//连接信息
type ConnInfo struct {
	Conn ziface.IConnection //所在连接
}

//单个道具信息
type SingleProp struct {
	ID  int64  //道具ID
	Num uint64 //数量
}
