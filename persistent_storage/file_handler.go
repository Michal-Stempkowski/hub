package persistent_storage

import (
	"hub/framework"
	"io/ioutil"
)

func ReadUserFile(name string) (res string, err error) {
	var data []byte
	data, err = ioutil.ReadFile(framework.GetUserDataPath(name))
	if err == nil {
		res = string(data)
	}
	return
}
