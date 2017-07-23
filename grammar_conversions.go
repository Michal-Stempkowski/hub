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

func generalizedCalculateFloatAndConvertToInt(
	expr FloatExpresion, rounding_function func(float64) float64) (result int64, err error) {
	var float_val float64
	float_val, err = expr.CalculateFloat()
	if err == nil {
		float_val = rounding_function(float_val)
		if float_val >= math.MaxInt64 || float_val <= math.MinInt64 {
			err = fmt.Errorf("Number %g to big to be represented as int")
		} else {
			result = int64(float_val)
		}
	}

	return
}

func round(val float64) float64 {
	return math.Floor(val + 0.5)
}

type FloatToIntConversion struct {
	*grammarElementImpl
	arg FloatExpresion
}

func (f *FloatToIntConversion) CalculateInt() (int64, error) {
	return generalizedCalculateFloatAndConvertToInt(f.arg, round)
}

func NewFloatToIntConversion(name string, arg FloatExpresion) *FloatToIntConversion {
	return &FloatToIntConversion{
		&grammarElementImpl{name, sheet_logic_types.FloatToIntConversion},
		arg}
}

type FloatToIntRoundDownConversion struct {
	*grammarElementImpl
	arg FloatExpresion
}

func (f *FloatToIntRoundDownConversion) CalculateInt() (result int64, err error) {
	return generalizedCalculateFloatAndConvertToInt(f.arg, math.Floor)
}

func NewFloatToIntRoundDownConversion(
	name string, arg FloatExpresion) *FloatToIntRoundDownConversion {
	return &FloatToIntRoundDownConversion{
		&grammarElementImpl{name, sheet_logic_types.FloatToIntRoundDownConversion},
		arg}
}

type FloatToIntRoundUpConversion struct {
	*grammarElementImpl
	arg FloatExpresion
}

func (f *FloatToIntRoundUpConversion) CalculateInt() (result int64, err error) {
	return generalizedCalculateFloatAndConvertToInt(f.arg, math.Ceil)
}

func NewFloatToIntRoundUpConversion(
	name string, arg FloatExpresion) *FloatToIntRoundUpConversion {
	return &FloatToIntRoundUpConversion{
		&grammarElementImpl{name, sheet_logic_types.FloatToIntRoundUpConversion},
		arg}
}
