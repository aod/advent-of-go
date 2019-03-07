package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/aoktayd/adventofgode/internal/error"
	"github.com/aoktayd/adventofgode/internal/input"
)

type claim struct {
	ID     int `json:"id"`
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type pos struct {
	x, y int
}

func main() {
	var claims []claim

	byteValue, err := ioutil.ReadFile(input.Path(input.File{
		Year:     2018,
		Day:      3,
		FileType: "json",
	}))
	error.Check(err)
	json.Unmarshal(byteValue, &claims)

	overlap := make(map[pos][]claim)
	claimIDs := make(map[int]struct{})

	for _, claim := range claims {
		claimIDs[claim.ID] = struct{}{}
		for y := claim.Y; y < claim.Y+claim.Height; y++ {
			for x := claim.X; x < claim.X+claim.Width; x++ {
				overlapPosition := pos{x, y}
				overlappingClaims := overlap[overlapPosition]
				overlap[overlapPosition] = append(overlappingClaims, claim)
			}
		}
	}

	totalInchesOverlap := 0

	for _, overlappingClaims := range overlap {
		if len(overlappingClaims) >= 2 {
			totalInchesOverlap++
			for _, claim := range overlappingClaims {
				delete(claimIDs, claim.ID)
			}
		}
	}

	fmt.Println("Part 1: ", totalInchesOverlap)
	fmt.Println("Part 2: ", claimIDs)
}
