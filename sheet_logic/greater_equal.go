package sheet_logic

import (
	"hub/framework"
	"hub/sheet_logic/sheet_logic_types"
)

type IntGreaterEqual IntComparator

func NewIntGreaterEqual(name string) *IntGreaterEqual {
	tmp := NewIntComparator(
		name,
		sheet_logic_types.IntGreaterEqual,
		func(a int64, b int64) bool { return a >= b })
	return (*IntGreaterEqual)(tmp)
}

type FloatGreaterEqual FloatComparator

func NewFloatGreaterEqual(name string) *FloatGreaterEqual {
	tmp := NewFloatComparator(
		name,
		sheet_logic_types.FloatGreaterEqual,
		framework.FloatGe)
	return (*FloatGreaterEqual)(tmp)
}

type StringGreaterEqual StringComparator

func NewStringGreaterEqual(name string) *StringGreaterEqual {
	tmp := NewStringComparator(
		name,
		sheet_logic_types.StringGreaterEqual,
		func(a string, b string) bool { return a >= b })
	return (*StringGreaterEqual)(tmp)
}
