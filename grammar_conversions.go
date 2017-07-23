package sheet_logic

import (
	"fmt"
	"hub/sheet_logic_types"
	"math"
	"strconv"
	"strings"
)

type IntToStringConversion struct {
	*grammarElementImpl
	arg IntExpresion
}

func (i *IntToStringConversion) CalculateString() (result string, err error) {
	var intVal int64
	intVal, err = i.arg.CalculateInt()
	if err == nil {
		result = strconv.FormatInt(intVal, 10)
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
	var intVal int64
	intVal, err = i.arg.CalculateInt()
	if err == nil {
		result = float64(intVal)
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
	var floatVal float64
	floatVal, err = expr.CalculateFloat()
	if err == nil {
		floatVal = rounding_function(floatVal)
		if floatVal >= math.MaxInt64 || floatVal <= math.MinInt64 {
			err = fmt.Errorf("Number %g to big to be represented as int")
		} else {
			result = int64(floatVal)
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

type StringToIntConversion struct {
	*grammarElementImpl
	arg StringExpresion
}

func (s *StringToIntConversion) CalculateInt() (result int64, err error) {
	var stringVal string
	stringVal, err = s.arg.CalculateString()
	if err == nil {
		result, err = strconv.ParseInt(stringVal, 10, 64)
	}

	return
}

func NewStringToIntConversion(name string, arg StringExpresion) *StringToIntConversion {
	return &StringToIntConversion{
		&grammarElementImpl{name, sheet_logic_types.StringToIntConversion},
		arg}
}

type StringToFloatConversion struct {
	*grammarElementImpl
	arg StringExpresion
}

func (s *StringToFloatConversion) CalculateFloat() (result float64, err error) {
	var stringVal string
	stringVal, err = s.arg.CalculateString()
	if err == nil {
		result, err = strconv.ParseFloat(stringVal, 64)
	}

	return
}

func NewStringToFloatConversion(name string, arg StringExpresion) *StringToFloatConversion {
	return &StringToFloatConversion{
		&grammarElementImpl{name, sheet_logic_types.StringToFloatConversion},
		arg}
}

type FloatToStringConversion struct {
	*grammarElementImpl
	arg       FloatExpresion
	Precision int64
}

func (f *FloatToStringConversion) CalculateString() (result string, err error) {
	var floatVal float64
	floatVal, err = f.arg.CalculateFloat()
	if f.Precision < 0 {
		err = fmt.Errorf("Negative precision not supported: %d", f.Precision)
	}

	if err == nil {
		result = strconv.FormatFloat(floatVal, 'f', -1, 64)
		dotIndex := int64(strings.Index(result, "."))
		if shift := dotIndex + f.Precision + 1; shift < int64(len(result)) {
			result = result[:shift]
		}
	}
	return
}

func NewFloatToStringConversion(
	name string, arg FloatExpresion, precision int64) *FloatToStringConversion {
	return &FloatToStringConversion{
		&grammarElementImpl{name, sheet_logic_types.FloatToStringConversion},
		arg,
		precision}
}
