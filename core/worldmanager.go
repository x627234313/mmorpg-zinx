package core

import "sync"

// 当前游戏的世界总管理模块
type WorldManager struct {
	// 当前地图AOI管理模块
	AoiManager *AOIManager

	// 当前全部的在线玩家集合
	Players map[uint32]*Player
	lock    sync.RWMutex
}

// 提供一个对外的世界管理模块的句柄（全局）
var WorldMgr *WorldManager

// 初始化方法
func init() {
	WorldMgr = &WorldManager{
		AoiManager: NewAOIManager(AOI_MIN_X, AOI_MAX_X, AOI_COUNTS_X, AOI_MIN_Y, AOI_MAX_Y, AOI_COUNTS_Y),
		Players:    make(map[uint32]*Player),
	}
}

// 添加一个玩家
func (wm *WorldManager) AddPlayer(p *Player) {
	wm.lock.Lock()
	wm.Players[p.PId] = p
	wm.lock.Unlock()

	// 将玩家添加到地图中
	wm.AoiManager.AddPlayerIdToGridByPos(int(p.PId), p.X, p.Z)
}

// 删除一个玩家
func (wm *WorldManager) RemovePlayer(pid uint32) {
	player := wm.Players[pid]

	// 从地图中删除
	wm.AoiManager.RemovePlayerIdFromGridByPos(int(player.PId), player.X, player.Z)

	// 从世界管理模块中删除
	wm.lock.Lock()
	delete(wm.Players, pid)
	wm.lock.Unlock()
}

// 获取一个玩家
func (wm *WorldManager) GetPlayer(pid uint32) *Player {
	wm.lock.RLock()
	defer wm.lock.RUnlock()

	return wm.Players[pid]
}

// 获取全部玩家
func (wm *WorldManager) GetAllPlayers() []*Player {
	wm.lock.RLock()
	defer wm.lock.RUnlock()

	players := make([]*Player, 0)

	for _, p := range wm.Players {
		players = append(players, p)
	}

	return players
}
