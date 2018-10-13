package utils

import (
	"path/filepath"
	"runtime"
)

// RelativePath Returns a relative path from the function caller
func RelativePath(fileName string) string {
	_, file, _, _ := runtime.Caller(1)
	return filepath.Join(filepath.Dir(file), fileName)
}
