package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
)

type And BoolComparator

func NewAnd(name string) *And {
	tmp := NewBoolComparator(
		name,
		sheet_logic_types.And,
		func(a bool, b bool) bool { return a && b })
	return (*And)(tmp)
}
