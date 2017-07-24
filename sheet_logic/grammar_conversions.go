package sheet_logic

import (
	"fmt"
	"hub/framework"
	"hub/sheet_logic/sheet_logic_types"
	"math"
	"strconv"
	"strings"
)

type IntToStringConversion struct {
	GrammarElement
	UnaryOperationInt
}

func (i *IntToStringConversion) CalculateString() (result string, err error) {
	var intVal int64

	if intVal, err = i.GetArg().CalculateInt(); err == nil {
		result = strconv.FormatInt(intVal, 10)
	}

	return
}

func NewIntToStringConversion(name string) *IntToStringConversion {
	return &IntToStringConversion{
		&grammarElementImpl{name, sheet_logic_types.IntToStringConversion},
		DefaultUnaryOperationIntImpl()}
}

type IntToFloatConversion struct {
	GrammarElement
	UnaryOperationInt
}

func (i *IntToFloatConversion) CalculateFloat() (result float64, err error) {
	var intVal int64

	if intVal, err = i.GetArg().CalculateInt(); err == nil {
		result = float64(intVal)
	}

	return
}

func NewIntToFloatConversion(name string) *IntToFloatConversion {
	return &IntToFloatConversion{
		&grammarElementImpl{name, sheet_logic_types.IntToFloatConversion},
		DefaultUnaryOperationIntImpl()}
}

func generalizedCalculateFloatAndConvertToInt(
	expr FloatExpresion, rounding_function func(float64) float64) (result int64, err error) {
	var floatVal float64

	if floatVal, err = expr.CalculateFloat(); err == nil {
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
	GrammarElement
	UnaryOperationFloat
}

func (f *FloatToIntConversion) CalculateInt() (int64, error) {
	return generalizedCalculateFloatAndConvertToInt(f.GetArg(), round)
}

func NewFloatToIntConversion(name string) *FloatToIntConversion {
	return &FloatToIntConversion{
		&grammarElementImpl{name, sheet_logic_types.FloatToIntConversion},
		DefaultUnaryOperationFloatImpl()}
}

type FloatToIntRoundDownConversion struct {
	GrammarElement
	UnaryOperationFloat
}

func (f *FloatToIntRoundDownConversion) CalculateInt() (result int64, err error) {
	return generalizedCalculateFloatAndConvertToInt(f.GetArg(), math.Floor)
}

func NewFloatToIntRoundDownConversion(name string) *FloatToIntRoundDownConversion {
	return &FloatToIntRoundDownConversion{
		&grammarElementImpl{name, sheet_logic_types.FloatToIntRoundDownConversion},
		DefaultUnaryOperationFloatImpl()}
}

type FloatToIntRoundUpConversion struct {
	GrammarElement
	UnaryOperationFloat
}

func (f *FloatToIntRoundUpConversion) CalculateInt() (result int64, err error) {
	return generalizedCalculateFloatAndConvertToInt(f.GetArg(), math.Ceil)
}

func NewFloatToIntRoundUpConversion(name string) *FloatToIntRoundUpConversion {
	return &FloatToIntRoundUpConversion{
		&grammarElementImpl{name, sheet_logic_types.FloatToIntRoundUpConversion},
		DefaultUnaryOperationFloatImpl()}
}

type StringToIntConversion struct {
	GrammarElement
	arg StringExpresion
}

func (s *StringToIntConversion) CalculateInt() (result int64, err error) {
	var stringVal string

	if stringVal, err = s.arg.CalculateString(); err == nil {
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
	GrammarElement
	arg StringExpresion
}

func (s *StringToFloatConversion) CalculateFloat() (result float64, err error) {
	var stringVal string

	if stringVal, err = s.arg.CalculateString(); err == nil {
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
	GrammarElement
	UnaryOperationFloat
	Precision int64
}

func (f *FloatToStringConversion) CalculateString() (result string, err error) {
	var floatVal float64
	floatVal, err = f.GetArg().CalculateFloat()

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

func NewFloatToStringConversion(name string) *FloatToStringConversion {
	return &FloatToStringConversion{
		&grammarElementImpl{name, sheet_logic_types.FloatToStringConversion},
		DefaultUnaryOperationFloatImpl(),
		-1}
}

type IntToBoolConversion struct {
	GrammarElement
	UnaryOperationInt
}

func (i *IntToBoolConversion) CalculateBool() (result bool, err error) {
	var intVal int64

	if intVal, err = i.GetArg().CalculateInt(); err == nil {
		if intVal == 0 {
			result = false
		} else {
			result = true
		}
	}

	return
}

func NewIntToBoolConversion(name string) *IntToBoolConversion {
	return &IntToBoolConversion{
		&grammarElementImpl{name, sheet_logic_types.IntToBoolConversion},
		DefaultUnaryOperationIntImpl()}
}

type FloatToBoolConversion struct {
	GrammarElement
	UnaryOperationFloat
}

func (f *FloatToBoolConversion) CalculateBool() (result bool, err error) {
	var floatVal float64

	if floatVal, err = f.GetArg().CalculateFloat(); err == nil {
		result = math.IsInf(floatVal, 1) ||
			math.IsInf(floatVal, -1) ||
			math.IsNaN(floatVal) ||
			math.Abs(floatVal) >= framework.FloatPrecision
	}

	return
}

func NewFloatToBoolConversion(name string) *FloatToBoolConversion {
	return &FloatToBoolConversion{
		&grammarElementImpl{name, sheet_logic_types.FloatToBoolConversion},
		DefaultUnaryOperationFloatImpl()}
}

type StringToBoolConversion struct {
	GrammarElement
	arg StringExpresion
}

func (s *StringToBoolConversion) CalculateBool() (result bool, err error) {
	var stringVal string

	if stringVal, err = s.arg.CalculateString(); err == nil {
		switch stringVal {
		case "true":
			result = true
		case "false":
			result = false
		default:
			err = fmt.Errorf(
				"StringToBoolConversion.CalculateBool failed for '%s'",
				stringVal)
		}
	}

	return
}

func NewStringToBoolConversion(name string, arg StringExpresion) *StringToBoolConversion {
	return &StringToBoolConversion{
		&grammarElementImpl{name, sheet_logic_types.StringToBoolConversion},
		arg}
}

type BoolToIntConversion struct {
	GrammarElement
	arg BoolExpresion
}

func (b *BoolToIntConversion) CalculateInt() (result int64, err error) {
	var boolVal bool

	if boolVal, err = b.arg.CalculateBool(); err == nil {
		if boolVal {
			result = 1
		} else {
			result = 0
		}
	}

	return
}

func NewBoolToIntConversion(name string, arg BoolExpresion) *BoolToIntConversion {
	return &BoolToIntConversion{
		&grammarElementImpl{name, sheet_logic_types.BoolToIntConversion},
		arg}
}

type BoolToFloatConversion struct {
	GrammarElement
	arg BoolExpresion
}

func (b *BoolToFloatConversion) CalculateFloat() (result float64, err error) {
	var boolVal bool

	if boolVal, err = b.arg.CalculateBool(); err == nil {
		if boolVal {
			result = 1.
		} else {
			result = 0.
		}
	}

	return
}

func NewBoolToFloatConversion(name string, arg BoolExpresion) *BoolToFloatConversion {
	return &BoolToFloatConversion{
		&grammarElementImpl{name, sheet_logic_types.BoolToFloatConversion},
		arg}
}

type BoolToStringConversion struct {
	GrammarElement
	arg BoolExpresion
}

func (b *BoolToStringConversion) CalculateString() (result string, err error) {
	var boolVal bool

	if boolVal, err = b.arg.CalculateBool(); err == nil {
		if boolVal {
			result = "true"
		} else {
			result = "false"
		}
	}

	return
}

func NewBoolToStringConversion(name string, arg BoolExpresion) *BoolToStringConversion {
	return &BoolToStringConversion{
		&grammarElementImpl{name, sheet_logic_types.BoolToStringConversion},
		arg}
}
