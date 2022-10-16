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
