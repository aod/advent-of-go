package input

import (
	"fmt"
	"go/build"
	"path/filepath"
	"strconv"

	"github.com/aoktayd/adventofgode/configs"
)

const inputFileName = "input"

// File : defines a puzzle input file
type File struct {
	Year, Day int
	FileType  string
}

// Path : returns the file path for the puzzle input
func Path(inputFile File) string {
	return filepath.Join(
		build.Default.GOPATH,
		configs.ProjectPath,
		strconv.Itoa(inputFile.Year),
		fmt.Sprintf("%02v", strconv.Itoa(inputFile.Day)),
		inputFileName+"."+inputFile.FileType,
	)
}
