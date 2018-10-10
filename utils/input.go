package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

// PuzzleInput Get the input of the year/day
func PuzzleInput(year uint16, day byte) []byte {
	_, file, _, _ := runtime.Caller(0)

	challengePath := fmt.Sprintf("%d/%02d/input.txt", year, day)
	path := filepath.Join(file, fmt.Sprintf("../../%s", challengePath))

	if len(os.Args) > 1 {
		return []byte(os.Args[1])
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return data
}
