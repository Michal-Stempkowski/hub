package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
)

type IntLesser IntComparator

func NewIntLesser(name string) *IntLesser {
	tmp := NewIntComparator(
		name,
		sheet_logic_types.IntLesser,
		func(a int64, b int64) bool { return a < b })
	return (*IntLesser)(tmp)
}
