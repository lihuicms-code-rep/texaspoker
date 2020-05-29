package apis

import (
	"github.com/golang/protobuf/proto"
	"github.com/lihuicms-code-rep/texaspoker/core"
	"github.com/lihuicms-code-rep/texaspoker/defination"
	"github.com/lihuicms-code-rep/texaspoker/log"
	"github.com/lihuicms-code-rep/texaspoker/pb"
	"github.com/lihuicms-code-rep/zinx/ziface"
	"github.com/lihuicms-code-rep/zinx/znet"
)

//玩家入桌业务,为客户端与服务器第一条交互
type PlayerEnterTable struct {
	znet.BaseRouter
}

//入桌流程
//1.load玩家数据
//2.查询可进入的桌子
//3.玩家入桌
//4.通知客户端入桌成功等信息
//5.广播玩家入桌
func (p *PlayerEnterTable) Handle(request ziface.IRequest) {
	msg := &pb.CS_Req_EnterTable{}
	err := proto.Unmarshal(request.GetData(), msg)
	if err != nil {
		log.Logic.Errorf("unmarshal request enter table error:%+v", err)
		return
	}


	//1.load玩家数据
	player := core.NewPlayer(msg.GetPid(), request.GetConnection())
	if player == nil {
		log.Logic.Errorf("new player:%d error", msg.GetPid())
		return
	} else {
		//绑定玩家所在连接
		conn := request.GetConnection()
		conn.SetProperty("pid", msg.GetPid())
	}

	//2.进入具体桌子
	table, err := core.TableMgrObj.GetTableByType(uint8(msg.GetTLevel()), uint8(msg.GetBLevel()))
	if err != nil {
		log.Logic.Errorf("get table by type error:%+v", err)
		player.SendMsg(pb.GameMessage_Ops_Rsp_EnterTable, &pb.SC_Rsp_EnterTable{
			Pid:msg.GetPid(),
			TableID:0,
			ChairID:0,
			Code:int32(defination.ErrorCodeNotCorrectTable),
		})
	}

	//3.玩家进入桌子
	

}
