package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/aoktayd/adventofcode-go/internal/error"
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

type claimGrid map[pos][]claim

func main() {
	var claims []claim

	fileContents, err := ioutil.ReadFile("input.json")
	error.Check(err)
	json.Unmarshal(fileContents, &claims)

	grid := make(claimGrid)
	claimIDs := make(map[int]struct{})
	totalInchesOverlap := 0

	for _, claim := range claims {
		claimIDs[claim.ID] = struct{}{}
		assignClaimToGrid(grid, claim)
	}

	for _, assignedClaims := range grid {
		if len(assignedClaims) < 2 {
			continue
		}
		totalInchesOverlap++
		for _, claim := range assignedClaims {
			delete(claimIDs, claim.ID)
		}
	}

	fmt.Println("Part 1: ", totalInchesOverlap)
	fmt.Println("Part 2: ", claimIDs)
}

func assignClaimToGrid(grid claimGrid, claim claim) {
	for y := claim.Y; y < claim.Y+claim.Height; y++ {
		for x := claim.X; x < claim.X+claim.Width; x++ {
			claimPosition := pos{x, y}
			assignedClaims := grid[claimPosition]
			grid[claimPosition] = append(assignedClaims, claim)
		}
	}
}
