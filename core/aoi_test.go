package core

import (
	"fmt"
	"testing"
)

// AOI管理模块的单元测试
func TestNewAOIManager(t *testing.T) {
	aoiMgr := NewAOIManager(0, 250, 5, 0, 250, 5)

	fmt.Println(aoiMgr)
}

// 获取当前格子周围的格子 功能测试
func TestGetSurroundGridByGid(t *testing.T) {
	aoiMgr := NewAOIManager(0, 250, 5, 0, 250, 5)

	for gid, _ := range aoiMgr.grids {
		sudoku := aoiMgr.GetSurroundGridsByGid(gid)
		fmt.Println("gid =", gid, "grids len =", len(sudoku))

		gIDs := make([]int, 0, len(sudoku))
		for _, grid := range sudoku {
			gIDs = append(gIDs, grid.gid)
		}
		fmt.Println("surround grid IDs are ", gIDs)
	}
}
