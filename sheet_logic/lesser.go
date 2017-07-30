package sheet_logic

import (
	"hub/framework"
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

type FloatLesser FloatComparator

func NewFloatLesser(name string) *FloatLesser {
	tmp := NewFloatComparator(
		name,
		sheet_logic_types.FloatLesser,
		framework.FloatLs)
	return (*FloatLesser)(tmp)
}

type StringLesser StringComparator

func NewStringLesser(name string) *StringLesser {
	tmp := NewStringComparator(
		name,
		sheet_logic_types.StringLesser,
		func(a string, b string) bool { return a < b })
	return (*StringLesser)(tmp)
}
