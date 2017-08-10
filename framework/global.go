package framework

import (
	"math"
	"os"
	"path/filepath"
)

const floatPrecision = 0.000001

func Round(f float64) float64 {
	return math.Floor(f + .5)
}

func FloatEq(a, b float64) bool {
	return math.Abs(a-b) < floatPrecision
}

func FloatLs(a, b float64) bool {
	return a+floatPrecision <= b
}

func FloatLe(a, b float64) bool {
	return FloatLs(a, b) || FloatEq(a, b)
}

func FloatGt(a, b float64) bool {
	return a >= b+floatPrecision
}

func FloatGe(a, b float64) bool {
	return FloatGt(a, b) || FloatEq(a, b)
}

var executableFinder = os.Executable

func SetDummyExecutableFinder() {
	executableFinder = func() (string, error) {
		return `C:\Users\Dell\go\src\hub\bin\run.exe`, nil
	}
}

func ResetRealExecutableFinder() {
	executableFinder = os.Executable
}

func GetBinaryPath() string {
	ex, err := executableFinder()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}

func GetRootPath() string {
	return filepath.Dir(GetBinaryPath())
}

func GetHtmlTemplatePath(name string) string {
	return filepath.Join(GetRootPath(), "html", name)
}

func GetUserDataPath(name string) string {
	return filepath.Join(GetRootPath(), "user_data", name)
}
