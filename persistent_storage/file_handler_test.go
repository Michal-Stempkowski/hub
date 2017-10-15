package persistent_storage

import (
	"hub/framework"
	"hub/persistent_storage/file_operation_status"
	"io/ioutil"
	"os"
	"testing"
)

const (
	readme           = "README"
	file1            = "dummy_file.ext"
	file2            = "dummy_file.ext2"
	notExistingFile  = "not_existing_file.ext"
	notExistingFile2 = "not_existing_file2.ext"
	file1Content     = "Content of file 1"
	file2Content     = "Content of file 2"
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

func TestDoesFileExist(t *testing.T) {
	runInFileHandlingTestContext(func() {
		assertFileExists(t, file1, "TestDoesFileExist")
		assertFileDoesNotExists(t, notExistingFile, "TestDoesFileExist")
	})
}

func TestSafeRename(t *testing.T) {
	runInFileHandlingTestContext(func() {
		filePath1 := framework.GetUserDataPath(file1)
		filePath2 := framework.GetUserDataPath(file2)
		filePathNew := framework.GetUserDataPath(notExistingFile)
		filePathNew2 := framework.GetUserDataPath(notExistingFile2)

		err := SafeRename(filePathNew, filePathNew2)
		assertFileOperationStatus(
			t, err, file_operation_status.DoesNotExist,
			"TestSafeRename: should be impossible to rename not existing file!")

		err = SafeRename(filePath1, filePath1)
		assertFileOperationStatus(
			t, err, file_operation_status.NameIsTheSame,
			"TestSafeRename: renaming to the same name should not be possible!")

		err = SafeRename(filePath1, filePath2)
		assertFileOperationStatus(
			t, err, file_operation_status.AlreadyExists,
			"TestSafeRename: should not rename if file already exists!")

		// Given:
		assertFileExists(t, file1, "TestSafeRename(Given)")
		assertFileDoesNotExists(t, notExistingFile, "TestSafeRename(Given)")
		// When:
		err = SafeRename(filePath1, filePathNew)
		if err != nil {
			t.Errorf("TestSafeRename(When): Rename should work!")
		} else {
			// Important so runInFileHandlingTestContext cleans proper resources!
			defer SafeRename(filePathNew, filePath1)

			// Then:
			assertFileDoesNotExists(t, file1, "TestSafeRename(Then)")
			assertFileExists(t, notExistingFile, "TestSafeRename(Then)")
		}

	})
}

func assertFileExists(t *testing.T, file, testcaseName string) {
	if !DoesFileExist(framework.GetUserDataPath(file)) {
		t.Errorf("%v: %v should exist", testcaseName, file)
	}
}

func assertFileDoesNotExists(t *testing.T, file, testcaseName string) {
	if DoesFileExist(framework.GetUserDataPath(file)) {
		t.Errorf("%v: %v should not exist", testcaseName, file)
	}
}

func assertFileOperationStatus(t *testing.T, err error, expectedStatus, message string) {
	if err == nil || err.Error() != expectedStatus {
		t.Errorf(message)
	}
}

func TestReadUserFile(t *testing.T) {
	runInFileHandlingTestContext(func() {
		data, err := ReadUserFile(readme)
		if err != nil || data != "This is where application stores its files" {
			t.Errorf("TestReadUserFile: improper data: '%v' or error: %v", data, err)
		}

		data, err = ReadUserFile("SomeNotExistingFile")
		if err == nil {
			t.Errorf("TestReadUserFile: This file should not exist!")
		}
	})
}

func expectFiles(
	t *testing.T, filesExpected, filesReceived []string, err error, testcaseName string) {
	if err != nil {
		t.Errorf("TestFilterFiles: On valid path error has occurred: %v", err)
	}

	receivedSet := make(map[string]bool)
	for _, file := range filesReceived {
		receivedSet[file] = true
	}

	for _, file := range filesExpected {
		if _, ok := receivedSet[file]; !ok {
			t.Errorf("%v: file %v not found while expected!", testcaseName, file)
		}
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
		expectFiles(
			t,
			[]string{file1, file2},
			files,
			err,
			"TestFilterFiles(no fillter)")

		files, err = FilterFiles(user_data_path, matchAllFiles)
		expectFiles(
			t,
			[]string{file1, file2},
			files,
			err,
			"TestFilterFiles(match all files)")

		files, err = FilterFiles(user_data_path, matchNoFiles)
		expectFiles(
			t,
			[]string{},
			files,
			err,
			"TestFilterFiles(match no files)")

		files, err = FilterFiles(user_data_path, GetExtensionFileInfoMatcher("ext"))
		expectFiles(
			t,
			[]string{file1},
			files,
			err,
			"TestFilterFiles(match *.ext files)")
	})
}

func TestHasExtension(t *testing.T) {
	if !HasExtension(file1, "ext") {
		t.Errorf("TestHasExtension: should have that extension!")
	}

	if HasExtension(file1, "xt") {
		t.Errorf("TestHasExtension: should not have that extension!")
	}
}
