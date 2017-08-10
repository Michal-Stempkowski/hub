package persistent_storage

import (
	"hub/framework"
	"testing"
)

func TestReadUserFile(t *testing.T) {
	framework.SetDummyExecutableFinder()
	defer framework.ResetRealExecutableFinder()
	data, err := ReadUserFile("README")
	if err != nil || data != "This is where application stores its files" {
		t.Errorf("TestReadUserFile: improper data: '%v' or error: %v", data, err)
	}
}
