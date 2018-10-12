package utils

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
)

// RelFile Gets a file relative to the function caller
func RelFile(fileName string) []byte {
	_, file, _, _ := runtime.Caller(1)
	path := filepath.Join(filepath.Dir(file), fileName)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return data
}
