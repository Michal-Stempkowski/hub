package sheet_logic

import (
	"hub/sheet_logic_types"
	"strconv"
)

type IntToStringConversion struct {
	*grammarElementImpl
	arg *IntConstant
}

func (i *IntToStringConversion) CalculateString() (result string, err error) {
	var int_val int
	int_val, err = i.arg.CalculateInt()
	if err == nil {
		result = strconv.Itoa(int_val)
	}

	return
}

func NewIntToStringConversion(name string, arg *IntConstant) *IntToStringConversion {
	return &IntToStringConversion{
		&grammarElementImpl{name, sheet_logic_types.IntToStringConversion},
		arg}
}

type IntToFloatConversion struct {
	*grammarElementImpl
	arg *IntConstant
}

func (i *IntToFloatConversion) CalculateFloat() (result float64, err error) {
	var int_val int
	int_val, err = i.arg.CalculateInt()
	if err == nil {
		result = float64(int_val)
	}

	return
}

func NewIntToFloatConversion(name string, arg *IntConstant) *IntToFloatConversion {
	return &IntToFloatConversion{
		&grammarElementImpl{name, sheet_logic_types.IntToFloatConversion},
		arg}
}
