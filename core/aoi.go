package core

import "fmt"

// AOI 区块管理模块
type AOIManager struct {
	// 区域的左边界坐标
	minX int
	// 区域的右边界坐标
	maxX int
	// X坐标方向格子的数量
	countX int

	// 区域的上边界坐标
	minY int
	// 区域的下边界坐标
	maxY int
	// Y坐标方向格子的数量
	countY int

	// 区域中所有的格子对象
	grids map[int]*Grid
}

func NewAOIManager(minx, maxx, countx, miny, maxy, county int) *AOIManager {
	aoiMgr := &AOIManager{
		minX:   minx,
		maxX:   maxx,
		countX: countx,
		minY:   miny,
		maxY:   maxy,
		countY: county,
		grids:  make(map[int]*Grid),
	}

	// 给AOI初始化区域的所有格子进行编号和初始化
	for y := 0; y < county; y++ {
		for x := 0; x < countx; x++ {
			// 计算格子编号，根据x，y轴的格子数量
			gid := y*countx + x

			aoiMgr.grids[gid] = NewGrid(gid,
				aoiMgr.minX+x*aoiMgr.gridWidth(), aoiMgr.minX+(x+1)*aoiMgr.gridWidth(),
				aoiMgr.minY+y*aoiMgr.gridHeight(), aoiMgr.minY+(y+1)*aoiMgr.gridHeight())
		}
	}

	return aoiMgr
}

// 计算每个格子x轴方向的宽度
func (m *AOIManager) gridWidth() int {
	return (m.maxX - m.minX) / m.countX
}

// 计算每个格子y轴方向的高度
func (m *AOIManager) gridHeight() int {
	return (m.maxY - m.minY) / m.countY
}

func (m *AOIManager) String() string {
	// 打印AOIManager信息
	s := fmt.Sprintf("AOIManager:\nMinX=%d, MaxX=%d, CountX=%d, MinY=%d, MaxY=%d, CountY=%d\nGrids in AOIManager:\n",
		m.minX, m.maxX, m.countX, m.minY, m.maxY, m.countY)

	// 打印全部格子信息
	for _, grid := range m.grids {
		s += fmt.Sprintln(grid)
	}

	return s
}
