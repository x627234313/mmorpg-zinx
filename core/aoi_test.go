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
