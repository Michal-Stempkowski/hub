package sheet_logic

import (
	"fmt"
	"hub/sheet_logic_types"
	"math"
	"strconv"
)

type IntToStringConversion struct {
	*grammarElementImpl
	arg IntExpresion
}

func (i *IntToStringConversion) CalculateString() (result string, err error) {
	var int_val int64
	int_val, err = i.arg.CalculateInt()
	if err == nil {
		result = strconv.FormatInt(int_val, 10)
	}

	return
}

func NewIntToStringConversion(name string, arg IntExpresion) *IntToStringConversion {
	return &IntToStringConversion{
		&grammarElementImpl{name, sheet_logic_types.IntToStringConversion},
		arg}
}

type IntToFloatConversion struct {
	*grammarElementImpl
	arg IntExpresion
}

func (i *IntToFloatConversion) CalculateFloat() (result float64, err error) {
	var int_val int64
	int_val, err = i.arg.CalculateInt()
	if err == nil {
		result = float64(int_val)
	}

	return
}

func NewIntToFloatConversion(name string, arg IntExpresion) *IntToFloatConversion {
	return &IntToFloatConversion{
		&grammarElementImpl{name, sheet_logic_types.IntToFloatConversion},
		arg}
}

type FloatToIntConversion struct {
	*grammarElementImpl
	arg FloatExpresion
}

func (f *FloatToIntConversion) CalculateInt() (result int64, err error) {
	var float_val float64
	float_val, err = f.arg.CalculateFloat()
	if err == nil {
		float_val = math.Floor(float_val + 0.5)
		if float_val >= math.MaxInt64 || float_val <= math.MinInt64 {
			err = fmt.Errorf("Number %g to big to be represented as int")
		} else {
			result = int64(float_val)
		}
	}

	return
}

func NewFloatToIntConversion(name string, arg FloatExpresion) *FloatToIntConversion {
	return &FloatToIntConversion{
		&grammarElementImpl{name, sheet_logic_types.FloatToIntConversion},
		arg}
}
