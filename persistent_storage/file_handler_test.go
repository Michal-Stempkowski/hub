package persistent_storage

import (
	"hub/framework"
	"io/ioutil"
	"os"
	"testing"
)

const (
	readme       = "README"
	file1        = "dummy_file.ext"
	file2        = "dummy_file.ext2"
	file1Content = "Content of file 1"
	file2Content = "Content of file 2"
)

func runInFileHandlingTestContext(scenario func()) {
	framework.SetDummyExecutableFinder()
	defer framework.ResetRealExecutableFinder()

	filePath1 := framework.GetUserDataPath(file1)
	filePath2 := framework.GetUserDataPath(file2)

	ioutil.WriteFile(filePath1, []byte(file1Content), 0644)
	defer os.Remove(filePath1)

	ioutil.WriteFile(filePath2, []byte(file2Content), 0644)
	defer os.Remove(filePath2)

	scenario()
}

func TestReadUserFile(t *testing.T) {
	runInFileHandlingTestContext(func() {
		data, err := ReadUserFile("README")
		if err != nil || data != "This is where application stores its files" {
			t.Errorf("TestReadUserFile: improper data: '%v' or error: %v", data, err)
		}

		data, err = ReadUserFile("SomeNotExistingFile")
		if err == nil {
			t.Errorf("TestReadUserFile: This file should not exist!")
		}
	})
}

func expectFiles(t *testing.T, filesExpected, filesReceived []string, err error) {
	if err != nil {
		t.Errorf("TestFilterFiles: On valid path error has occurred: %v", err)
	}
	if len(filesReceived) != len(filesExpected) {
		t.Errorf(
			"TestFilterFiles: argument number differs (expected=%v, received=%v)",
			filesExpected,
			filesReceived)
	}
	for i := 0; i < len(filesExpected); i++ {
		expectFile(t, filesExpected[i], filesReceived[i])
	}
}

func expectFile(t *testing.T, fileExpected, fileReceived string) {
	if fileReceived != fileExpected {
		t.Errorf(
			"TestFilterFiles: First file, expected: %v, received: %v",
			fileExpected,
			fileReceived)
	}
}

func matchAllFiles(os.FileInfo) bool {
	return true
}

func matchNoFiles(os.FileInfo) bool {
	return false
}

func TestFilterFiles(t *testing.T) {
	runInFileHandlingTestContext(func() {
		files, err := FilterFiles("/some/not/existing/path", nil)
		if err == nil {
			t.Errorf("TestFilterFiles: On invalid path should return nil, received: %v", files)
		}

		user_data_path := framework.GetUserDataPath("")
		files, err = FilterFiles(user_data_path, nil)
		expectFiles(t, []string{readme, file1, file2}, files, err)

		files, err = FilterFiles(user_data_path, matchAllFiles)
		expectFiles(t, []string{readme, file1, file2}, files, err)

		files, err = FilterFiles(user_data_path, matchNoFiles)
		expectFiles(t, []string{}, files, err)

		files, err = FilterFiles(user_data_path, GetExtensionFileInfoMatcher(".ext"))
		expectFiles(t, []string{file1}, files, err)
	})
}
