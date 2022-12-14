package core

import "fmt"

// 定义地图边界常量
const (
	AOI_MIN_X    = 0
	AOI_MAX_X    = 600
	AOI_COUNTS_X = 50
	AOI_MIN_Y    = 0
	AOI_MAX_Y    = 500
	AOI_COUNTS_Y = 50
)

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

// 根据格子id-gid，获取周围格子信息
func (m *AOIManager) GetSurroundGridsByGid(gid int) []*Grid {
	// 声明周围九宫格切片
	var sudoku []*Grid

	// 判断当前id是否在AOI区域中
	if _, ok := m.grids[gid]; !ok {
		return nil
	}

	// 当前格子在AOI区域中，把当前格子加入到九宫格切片中
	sudoku = append(sudoku, m.grids[gid])

	// 计算当前格子所在的x轴的编号
	idx := gid % m.countX

	// 判断当前格子左边是否有格子
	if idx > 0 {
		// 左边有格子，把左边格子加入到九宫格切片中
		sudoku = append(sudoku, m.grids[gid-1])
	}
	// 判断当前格子右边是否有格子
	if idx < m.countX-1 {
		// 右边有格子，把右边格子加入到九宫格切片中
		sudoku = append(sudoku, m.grids[gid+1])
	}

	// 把x轴上的格子取出遍历，判断上下是否有格子
	// 声明一个切片，保存x轴上格子的id
	XgIDs := make([]int, 0, len(sudoku))
	for _, grid := range sudoku {
		XgIDs = append(XgIDs, grid.gid)
	}

	// 判断x轴上每个格子的上下是否有格子
	for _, xgid := range XgIDs {
		// 当前格子所在的y轴的编号
		idy := xgid / m.countY
		// xgid上边是否有格子，如果有格子，加入到九宫格切片中
		if idy > 0 {
			sudoku = append(sudoku, m.grids[xgid-m.countX])
		}
		// xgid下边是否有格子，如果有格子，加入到九宫格切片中
		if idy < m.countY-1 {
			sudoku = append(sudoku, m.grids[xgid+m.countX])
		}
	}

	return sudoku
}

// 根据玩家ID，获取周围九宫格内所有玩家ID
func (m *AOIManager) GetSurroundPlayerIDsByPos(x, y float32) (playerIDs []int) {
	// 首先根据x,y轴坐标，得到格子ID
	gid := m.GetGidByPos(x, y)

	// 再根据格子ID，得到九宫格ID
	sudoku := m.GetSurroundGridsByGid(gid)

	// 得到每个格子内的玩家ID
	for _, grid := range sudoku {
		playerIDs = append(playerIDs, grid.GetPlayerIDs()...)
	}

	return
}

// 根据 x，y 轴坐标得到格子ID
func (m *AOIManager) GetGidByPos(x, y float32) int {
	idx := (int(x) - m.minX) / m.gridWidth()
	idy := (int(y) - m.minY) / m.gridHeight()

	return idy*m.countX + idx
}

// 通过gid获取当前格子的全部playerId
func (m *AOIManager) GetPlayerIdByGid(gid int) (playerIDs []int) {
	playerIDs = m.grids[gid].GetPlayerIDs()
	return
}

// 添加一个playerId到一个格子中
func (m *AOIManager) AddPlayerIdToGrid(pid, gid int) {
	m.grids[gid].Add(pid)
}

// 移除一个playerId从一个格子中
func (m *AOIManager) RemovePlayerIdFromGrid(pid, gid int) {
	m.grids[gid].Remove(pid)
}

// 通过坐标把一个player添加到一个格子中
func (m *AOIManager) AddPlayerIdToGridByPos(pid int, x, y float32) {
	gid := m.GetGidByPos(x, y)
	m.grids[gid].Add(pid)
}

// 通过坐标把一个player从一个格子中移除
func (m *AOIManager) RemovePlayerIdFromGridByPos(pid int, x, y float32) {
	gid := m.GetGidByPos(x, y)
	m.grids[gid].Remove(pid)
}
