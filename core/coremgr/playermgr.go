package core

import "sync"

//玩家管理类
type PlayerManager struct {
	players map[uint32]*Player //当前在线玩家
	curNum  int                //当前玩家数
	pLock   sync.RWMutex       //player读写锁
}

//玩家管理对外使用对象
var PlayerMgrObj *PlayerManager

//手动初始化
func InitPlayerManager() {
	PlayerMgrObj = &PlayerManager{
		players: make(map[uint32]*Player),
		curNum:  0,
	}
}

//添加至玩家map
func (mgr *PlayerManager) AddPlayer(p *Player) {
	if p == nil {
		return
	}

	mgr.pLock.Lock()
	defer mgr.pLock.Unlock()

	if _, ok := mgr.players[p.BInfo.pid]; !ok {
		mgr.players[p.BInfo.pid] = p
		mgr.curNum += 1
	}
}

//从玩家map中移除
func (mgr *PlayerManager) RemovePlayer(p *Player) {
	if p == nil {
		return
	}

	mgr.pLock.Lock()
	defer mgr.pLock.Unlock()

	if _, ok := mgr.players[p.BInfo.pid]; ok {
		delete(mgr.players, p.BInfo.pid)
		mgr.curNum -= 1
	}
}

//根据玩家ID获取玩家信息
func (mgr *PlayerManager) GetPlayerByID(id uint32) *Player {
	mgr.pLock.RLock()
	defer mgr.pLock.RUnlock()

	p, ok := mgr.players[id]
	if !ok {
		return nil
	}

	return p
}

//获取当前在线玩家数
func (mgr *PlayerManager) GetOnlineNum() int {
	return mgr.curNum
}

//获取在线玩家列表
func (mgr *PlayerManager) GetOnlinePlayers() []*Player {
	mgr.pLock.RLock()
	defer mgr.pLock.RUnlock()

	players := make([]*Player, 0, mgr.curNum)
	for _, p := range mgr.players {
		players = append(players, p)
	}

	return players
}
