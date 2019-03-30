package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type position struct {
	x, y int
}

func (p *position) dist(to position) int {
	return int(math.Abs(float64(p.x-to.x)) + math.Abs(float64(p.y-to.y)))
}

func (p *position) closest(positions []*position) (*position, int, bool) {
	var closestPosition *position
	lowestDistance := math.MaxInt32
	distances := map[int]int{}

	for _, pos := range positions {
		dist := p.dist(*pos)
		distances[dist]++
		if dist < lowestDistance {
			lowestDistance = dist
			closestPosition = pos
		}
	}

	return closestPosition, lowestDistance, distances[lowestDistance] > 1
}

func main() {
	file, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(file)
	defer file.Close()

	var positions []*position
	for fileScanner.Scan() {
		var x, y int
		fmt.Sscanf(fileScanner.Text(), "%d, %d", &x, &y)
		positions = append(positions, &position{x: x, y: y})
	}

	minBoundingPosition, maxBoundingPosition := getBoundingPositions(positions)

	grid := map[position]*position{}
	area := map[*position][]*position{}
	for y := minBoundingPosition.y; y <= maxBoundingPosition.y; y++ {
		for x := minBoundingPosition.x; x <= maxBoundingPosition.x; x++ {
			pos := &position{x: x, y: y}
			closestPosition, _, hasTwoOrMoreClosestPositions := pos.closest(positions)

			if hasTwoOrMoreClosestPositions {
				continue
			}

			area[closestPosition] = append(area[closestPosition], pos)
			grid[*pos] = closestPosition
		}
	}

	// Loops through border to check infinite areas
	infiniteAreaPositions := map[*position]bool{}
	for y := minBoundingPosition.y; y <= maxBoundingPosition.y; y++ {
		isBetweenTopOrBottom := y > minBoundingPosition.y && y < maxBoundingPosition.y

		for x := minBoundingPosition.x; x <= maxBoundingPosition.x; x++ {
			pos := grid[position{x: x, y: y}]

			if _, exists := infiniteAreaPositions[pos]; pos != nil && !exists {
				infiniteAreaPositions[pos] = true
			}

			if isBetweenTopOrBottom {
				x += maxBoundingPosition.x - minBoundingPosition.x - 1
				continue
			}
		}
	}

	largestFiniteAreaSize := 0
	for pos, areaPositions := range area {
		if infiniteAreaPositions[pos] {
			continue
		}
		if areaSize := len(areaPositions); areaSize > largestFiniteAreaSize {
			largestFiniteAreaSize = areaSize
		}
	}

	safeAreaSize := 0
	for y := minBoundingPosition.y; y <= maxBoundingPosition.y; y++ {
		for x := minBoundingPosition.x; x <= maxBoundingPosition.x; x++ {
			gridPos := &position{x: x, y: y}
			sumDistance := 0
			for _, pos := range positions {
				sumDistance += gridPos.dist(*pos)
			}
			if sumDistance < 10000 {
				safeAreaSize++
			}
		}
	}

	fmt.Println("Part 1: ", largestFiniteAreaSize)
	fmt.Println("Part 2: ", safeAreaSize)
}

func getBoundingPositions(positions []*position) (position, position) {
	minBoundingPosition := position{math.MaxInt32, math.MaxInt32}
	maxBoundingPosition := position{0, 0}

	for _, pos := range positions {
		if pos.x > maxBoundingPosition.x {
			maxBoundingPosition.x = pos.x
		}
		if pos.y > maxBoundingPosition.y {
			maxBoundingPosition.y = pos.y
		}
		if pos.x < minBoundingPosition.x {
			minBoundingPosition.x = pos.x
		}
		if pos.y < minBoundingPosition.y {
			minBoundingPosition.y = pos.y
		}
	}

	return minBoundingPosition, maxBoundingPosition
}
