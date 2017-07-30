package sheet_logic

import (
	"hub/framework"
	"hub/sheet_logic/sheet_logic_types"
)

type IntGreater IntComparator

func NewIntGreater(name string) *IntGreater {
	tmp := NewIntComparator(
		name,
		sheet_logic_types.IntGreater,
		func(a int64, b int64) bool { return a > b })
	return (*IntGreater)(tmp)
}

type FloatGreater FloatComparator

func NewFloatGreater(name string) *FloatGreater {
	tmp := NewFloatComparator(
		name,
		sheet_logic_types.FloatGreater,
		framework.FloatGt)
	return (*FloatGreater)(tmp)
}

type StringGreater StringComparator

func NewStringGreater(name string) *StringGreater {
	tmp := NewStringComparator(
		name,
		sheet_logic_types.StringGreater,
		func(a string, b string) bool { return a > b })
	return (*StringGreater)(tmp)
}
