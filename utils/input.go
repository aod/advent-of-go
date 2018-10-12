package utils

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
)

// ReadRelFile Reads a file relative to the function caller
func ReadRelFile(fileName string) ([]byte, error) {
	_, file, _, _ := runtime.Caller(1)
	path := filepath.Join(filepath.Dir(file), fileName)

	data, err := ioutil.ReadFile(path)

	return data, err
}
