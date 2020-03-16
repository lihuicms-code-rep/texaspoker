package utils

import (
	"encoding/json"
	"io/ioutil"
)

//读取文件内容至对象
func ReadFile2Obj(filename string, obj interface{}) (bool, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return false, err
	}

	if err = json.Unmarshal(bytes, obj); err != nil {
		return false, err
	}

	return true, nil
}