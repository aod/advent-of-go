package input

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/aoktayd/adventofgode/configs"
	"github.com/aoktayd/adventofgode/internal/error"
)

// JSON represents the JSON file type
const JSON = ".json"

// File describes an AOC puzzle input
type File struct {
	Puzzle   Puzzle
	FileType string
}

const inputFileName = "input"

// Path returns the file path for the puzzle input
func Path(inputFile File) string {
	return filepath.Join(
		build.Default.GOPATH,
		configs.ProjectPath,
		strconv.Itoa(inputFile.Puzzle.Year),
		fmt.Sprintf("%02v", strconv.Itoa(inputFile.Puzzle.Day)),
		inputFileName+inputFile.FileType,
	)
}

// Read uses the ioutil pkg to return the contents of the input file
func Read(inputFile File) []byte {
	inputFilePath := Path(inputFile)
	byteValue, err := ioutil.ReadFile(inputFilePath)
	error.Check(err)
	return byteValue
}

// ReadJSON parses the input file as JSON and stores the result in the value pointed by v
func ReadJSON(aocPuzzle Puzzle, v interface{}) {
	inputFile := File{Puzzle: aocPuzzle, FileType: JSON}
	err := json.Unmarshal(Read(inputFile), &v)
	error.Check(err)
}

// Scanner does something
func Scanner(aocPuzzle Puzzle) (*bufio.Scanner, *os.File) {
	f, err := os.Open(Path(File{Puzzle: aocPuzzle, FileType: ".txt"}))
	error.Check(err)
	return bufio.NewScanner(f), f
}
