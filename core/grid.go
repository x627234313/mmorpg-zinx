package core

import (
	"fmt"
	"sync"
)

// AOI区块中每个格子对象
type Grid struct {
	// 格子ID
	gid int

	// 格子左边边界坐标
	minX int
	// 格子右边边界坐标
	maxX int
	// 格子上边边界坐标
	minY int
	// 格子下边边界坐标
	maxY int

	// 格子中玩家或者物体的ID的集合
	players map[int]bool

	playerLock sync.RWMutex
}

func NewGrid(id, minx, maxx, miny, maxy int) *Grid {
	return &Grid{
		gid:     id,
		minX:    minx,
		maxX:    maxx,
		minY:    miny,
		maxY:    maxy,
		players: make(map[int]bool),
	}
}

// 给一个格子添加玩家
func (g *Grid) Add(playerID int) {
	// 加写锁
	g.playerLock.Lock()
	defer g.playerLock.Unlock()

	g.players[playerID] = true
}

// 从格子中删除一个玩家
func (g *Grid) Remove(playerID int) {
	// 加写锁
	g.playerLock.Lock()
	defer g.playerLock.Unlock()

	delete(g.players, playerID)
}

// 得到当前格子中的所有玩家ID
func (g *Grid) GetPlayerIDs() (playerIDs []int) {
	// 加读锁
	g.playerLock.RLock()
	defer g.playerLock.RUnlock()

	for id, _ := range g.players {
		playerIDs = append(playerIDs, id)
	}

	return
}

// 调试 -- 打印格子信息
func (g *Grid) String() string {
	return fmt.Sprintf("Grid id=%d, MinX=%d, MaxX=%d, MinY=%d, MaxY=%d, playerIDs=%v",
		g.gid, g.minX, g.maxX, g.minY, g.maxY, g.players)
}
