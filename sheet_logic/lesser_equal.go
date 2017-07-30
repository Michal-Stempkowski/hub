package sheet_logic

import (
	"hub/framework"
	"hub/sheet_logic/sheet_logic_types"
)

type IntLesserEqual IntComparator

func NewIntLesserEqual(name string) *IntLesserEqual {
	tmp := NewIntComparator(
		name,
		sheet_logic_types.IntLesserEqual,
		func(a int64, b int64) bool { return a <= b })
	return (*IntLesserEqual)(tmp)
}

type FloatLesserEqual FloatComparator

func NewFloatLesserEqual(name string) *FloatLesserEqual {
	tmp := NewFloatComparator(
		name,
		sheet_logic_types.FloatLesserEqual,
		framework.FloatLe)
	return (*FloatLesserEqual)(tmp)
}

type StringLesserEqual StringComparator

func NewStringLesserEqual(name string) *StringLesserEqual {
	tmp := NewStringComparator(
		name,
		sheet_logic_types.StringLesserEqual,
		func(a string, b string) bool { return a <= b })
	return (*StringLesserEqual)(tmp)
}
