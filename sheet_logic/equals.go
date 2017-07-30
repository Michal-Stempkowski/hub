package sheet_logic

import (
	"hub/framework"
	"hub/sheet_logic/sheet_logic_types"
)

type IntEquals IntComparator

func NewIntEquals(name string) *IntEquals {
	tmp := NewIntComparator(
		name,
		sheet_logic_types.IntEquals,
		func(a int64, b int64) bool { return a == b })
	return (*IntEquals)(tmp)
}

type FloatEquals FloatComparator

func NewFloatEquals(name string) *FloatEquals {
	tmp := NewFloatComparator(
		name,
		sheet_logic_types.FloatEquals,
		framework.FloatEq)
	return (*FloatEquals)(tmp)
}

type BoolEquals BoolComparator

func NewBoolEquals(name string) *BoolEquals {
	tmp := NewBoolComparator(
		name,
		sheet_logic_types.BoolEquals,
		func(a bool, b bool) bool { return a == b })
	return (*BoolEquals)(tmp)
}

type StringEquals StringComparator

func NewStringEquals(name string) *StringEquals {
	tmp := NewStringComparator(
		name,
		sheet_logic_types.StringEquals,
		func(a string, b string) bool { return a == b })
	return (*StringEquals)(tmp)
}
