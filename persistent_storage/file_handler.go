package persistent_storage

import (
	"hub/framework"
	"io/ioutil"
	"os"
	"strings"
)

const RuleSetExtension = ".ruleset"

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

func GetExtensionFileInfoMatcher(ext string) func(f os.FileInfo) bool {
	return func(f os.FileInfo) bool {
		return strings.HasSuffix(f.Name(), ext)
	}
}
