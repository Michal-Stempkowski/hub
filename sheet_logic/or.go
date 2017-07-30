package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
)

type Or BoolComparator

func NewOr(name string) *Or {
	tmp := NewBoolComparator(
		name,
		sheet_logic_types.Or,
		func(a bool, b bool) bool { return a || b })
	return (*Or)(tmp)
}
