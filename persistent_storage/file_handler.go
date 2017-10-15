package persistent_storage

import (
	"fmt"
	"hub/framework"
	"hub/persistent_storage/file_operation_status"
	"io/ioutil"
	"os"
	"strings"
)

const RuleSetExtension = "ruleset"

func DoesFileExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func SafeRename(path, newPath string) error {
	switch {
	case !DoesFileExist(path):
		return fmt.Errorf(file_operation_status.DoesNotExist)
	case path == newPath:
		return fmt.Errorf(file_operation_status.NameIsTheSame)
	case DoesFileExist(newPath):
		return fmt.Errorf(file_operation_status.AlreadyExists)
	default:
		return os.Rename(path, newPath)
	}
}

func ReadUserFile(name string) (res string, err error) {
	var data []byte
	data, err = ioutil.ReadFile(framework.GetUserDataPath(name))
	if err == nil {
		res = string(data)
	}
	return
}

func FilterFiles(path string, filter func(os.FileInfo) bool) (result []string, err error) {
	var files []os.FileInfo
	files, err = ioutil.ReadDir(path)
	if err == nil {
		for _, fileInfo := range files {
			if filter == nil || filter(fileInfo) {
				result = append(result, fileInfo.Name())
			}
		}
	}
	return
}

func HasExtension(name, ext string) bool {
	return strings.HasSuffix(name, "."+ext)
}

func GetExtensionFileInfoMatcher(ext string) func(f os.FileInfo) bool {
	return func(f os.FileInfo) bool {
		return HasExtension(f.Name(), ext)
	}
}
