package core

import (
	"fmt"
	"github.com/lihuicms-code-rep/texaspoker/config"
	"github.com/lihuicms-code-rep/texaspoker/log"
	"github.com/pkg/errors"
	"sync"
)

//桌子管理类
//桌子分配,资源管理等
type TableManager struct {
	runningTables map[uint16]*Table   //按tableID映射当前在运行的桌子
	typeTables    map[string][]*Table //按桌子类型映射当前在运行的桌子
	//按桌子类型映射其桌子ID分配范围等,这里一个类型桌子的确定key为:tableLevel_blindLevel,如(1_2)表示六人桌的20/40级别
	idAlloc map[string]*tableIDAlloc
	tLock   sync.RWMutex
}

//不同类型的桌子ID动态分配
type tableIDAlloc struct {
	startID uint16 //起始ID
	endID   uint16 //截止ID
	nextID  uint16 //下一个要被使用的ID
	pause   bool   //分配完则先暂停分配,通知扩容
}

//桌子管理对外使用对象
var TableMgrObj *TableManager

//手动初始化
func InitTableManager() {
	TableMgrObj = &TableManager{
		runningTables: make(map[uint16]*Table),
		typeTables:    make(map[string][]*Table),
		idAlloc:       make(map[string]*tableIDAlloc),
	}

	for _, tc := range config.RoomConfig.TableList {
		key := fmt.Sprintf("%d_%d", tc.TableLevel, tc.BlindLevel)
		TableMgrObj.idAlloc[key] = &tableIDAlloc{
			startID: tc.StartID,
			endID:   tc.StartID + tc.Num - 1,
			nextID:  tc.StartID,
			pause:   false,
		}
	}
}

//加入运行桌子
func (mgr *TableManager) AddTable(table *Table) bool {
	if table == nil {
		return false
	}

	mgr.tLock.Lock()
	defer mgr.tLock.Unlock()

	if _, ok := mgr.runningTables[table.tableID]; ok {
		log.Logic.Warnf("this table:%d exist in running table", table.tableID)
		return false
	}
	mgr.runningTables[table.tableID] = table

	key := fmt.Sprintf("%d_%d", table.tableLevel, table.blindLevel)
	if _, ok := mgr.typeTables[key]; ok {
		log.Logic.Warnf("this table:%d exist in type table", table.tableID)
		return false
	}
	mgr.typeTables[key] = append(mgr.typeTables[key], table)

	if ac, ok := mgr.idAlloc[key]; ok {
		if ac.nextID == ac.endID {
			ac.pause = true
		} else {
			ac.nextID += 1
		}

		mgr.idAlloc[key] = ac //注意struct赋值的问题
	}

	return true
}

//移除停运桌子
func (mgr *TableManager) RemoveTable(table *Table) bool {
	if table == nil {
		return false
	}

	mgr.tLock.Lock()
	defer mgr.tLock.Unlock()

	if _, ok := mgr.runningTables[table.tableID]; !ok {
		log.Logic.Warnf("this table:%d not exist in running table", table.tableID)
		return false
	}

	delete(mgr.runningTables, table.tableID)

	//桌子人数类型+盲注类型唯一确定一种类型的桌子
	tType := fmt.Sprintf("%d_%d", table.tableLevel, table.blindLevel)
	if _, ok := mgr.typeTables[tType]; !ok {
		log.Logic.Warnf("this table:%d not exist in type table", table.tableID)
		return false
	}

	//查找该桌子在slice中的下标
	index := 0
	for i, t := range mgr.typeTables[tType] {
		if t == table {
			index = i
		}
	}

	//删除
	mgr.typeTables[tType] = append(mgr.typeTables[tType][:index], mgr.typeTables[tType][index+1:]...)
	return true
}

//根据桌子id获取桌子信息
func (mgr *TableManager) GetTableByID(id uint16) *Table {
	if table, ok := mgr.runningTables[id]; ok {
		return table
	}

	return nil
}

//查找一个类型中可进入玩家的桌子
func (mgr *TableManager) GetTableByType(tableLevel, blindLevel uint8) (*Table, error) {
	key := fmt.Sprintf("%d_%d", tableLevel, blindLevel)
	//此类型桌子还没有,需要创建一个
	tableList, ok := mgr.typeTables[key]
	if !ok {
		log.Logic.Warnf("this tableLevel:%d blindLevel:%d need to create", tableLevel, blindLevel)
		id := mgr.GetNextIDWhenCreateTable(tableLevel, blindLevel)
		if id == 0 {
			log.Logic.Warnf("this tableLevel:%d blindLevel:%d alloc pause", tableLevel, blindLevel)
			return nil, errors.Errorf("this tableLevel:%d blindLevel:%d alloc pause", tableLevel, blindLevel)
		}

		table := NewTable(config.GetTableCfgByID(id))
		mgr.AddTable(table)
		return table, nil
	}

	//此类型桌子还有,顺序查找第一个还可以进人的(这里做的比较简单)
	for _, t := range tableList {
		if t.curGamer < t.maxGamer {
			return t, nil
		}
	}

	return nil, nil
}

//当创建桌子时获取该类型桌子的nextID
//当无法再分配时,返回0
func (mgr *TableManager) GetNextIDWhenCreateTable(tableLevel, blindLevel uint8) uint16 {
	key := fmt.Sprintf("%d_%d", tableLevel, blindLevel)
	if !mgr.canCreateTable(tableLevel, blindLevel) {
		return 0
	}

	if at, ok := mgr.idAlloc[key]; ok {
		return at.nextID
	}

	return 0
}

//创建桌子前先询问该类型桌子是否还可以创建
//true:可以, false:不可以
func (mgr *TableManager) canCreateTable(tableLevel, blindLevel uint8) bool {
	key := fmt.Sprintf("%d_%d", tableLevel, blindLevel)
	if at, ok := mgr.idAlloc[key]; ok {
		return !at.pause //注意含义
	}

	return true
}
