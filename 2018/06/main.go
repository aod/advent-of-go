package main

import (
	"fmt"
	"math"

	"github.com/aoktayd/adventofgode/internal/input"
)

type coordinate struct {
	x, y int
}

func (c *coordinate) distance(to coordinate) int {
	return int(math.Abs(float64(c.x-to.x)) + math.Abs(float64(c.y-to.y)))
}

func (c *coordinate) closest(coordinates []coordinate) (*coordinate, bool) {
	var closestCoordinate *coordinate

	closestDistance := math.MaxInt32
	distances := make(map[int]int)

	for key, coordinate := range coordinates {
		dist := c.distance(coordinate)
		distances[dist]++
		if dist < closestDistance {
			closestDistance = dist
			closestCoordinate = &coordinates[key]
		}
	}

	return closestCoordinate, distances[closestDistance] >= 2
}

func main() {
	var coordinates []coordinate
	parseInput(&coordinates)

	maxCoordinate := maxBoundingCoordinate(coordinates)

	maxCoordinate.y++
	maxCoordinate.x++

	// Keeps track of the size of the locations (coordinates from the input a.k.a. input.txt)
	areas := make(map[coordinate]int)

	// This is for every position in the partial grid which is finite
	// whereas the actual grid extends infinitely in all directions.
	// Every coordinate on the grid points to the closest location.
	grid := make(map[coordinate]*coordinate)

	// Keeps track of what coordinates are closest to a location
	locationCoordinates := make(map[coordinate][]*coordinate)

	// Calculate the distances to every location of every coordinate on the grid
	for y := 0; y < maxCoordinate.y; y++ {
		for x := 0; x < maxCoordinate.x; x++ {
			gridCoordinate := coordinate{x, y}
			closestCoordinate, hasTwoOrMore := gridCoordinate.closest(coordinates)

			// If the coordinate on the grid has 2 or more locations with the same distance
			// the area size is not increased
			if hasTwoOrMore {
				continue
			}

			locationCoordinates[*closestCoordinate] = append(locationCoordinates[*closestCoordinate], &gridCoordinate)
			grid[gridCoordinate] = closestCoordinate
			areas[*closestCoordinate]++
		}
	}

	// Remove coordinates which extend infinitly by checking to which
	// main coordinate every grid coordinate points to at the border
	for y := 0; y < maxCoordinate.y; y++ {
		for x := 0; x < maxCoordinate.x; x++ {
			gridCoordinate := coordinate{x, y}

			// If we're not at the top or the bottom of the grid
			// only check the left and right side
			if y > 0 && y < maxCoordinate.y-1 {
				x += maxCoordinate.x - 2
			}

			coordinate, coordinateExists := grid[gridCoordinate]

			// fmt.Println("H", x, y, coordinate, coordinateExists)

			if !coordinateExists {
				continue
			}

			// fmt.Println("deleting", coordinate)

			delete(areas, *coordinate)
			delete(locationCoordinates, *coordinate)
		}
	}

	// The largest area size that isn't infinite
	biggestAreaSize := 0

	for _, areaSize := range areas {
		if areaSize > biggestAreaSize {
			biggestAreaSize = areaSize
		}
	}

	// This is the size of the region where the region is near as many coordinates as possible
	bestReagionSize := len(locationCoordinates)

	const maxDistance = 10000

	for _, lCoordinates := range locationCoordinates {
		for _, gridCoordinate := range lCoordinates {
			totalDistance := 0
			for _, location := range coordinates {
				totalDistance += gridCoordinate.distance(location)
			}
			if totalDistance < maxDistance {
				bestReagionSize++
			}
		}
	}

	// fmt.Println(areas)
	// fmt.Println(locationCoordinates)

	fmt.Println("Part 1: ", biggestAreaSize)
	// 42132 too high
	// 32268 Too low
	fmt.Println("Part 2: ", bestReagionSize)
}

func parseInput(coordinates *[]coordinate) {
	s, f := input.Scanner(input.Puzzle{Year: 2018, Day: 6})
	defer f.Close()

	for s.Scan() {
		var x, y int
		fmt.Sscanf(s.Text(), "%d, %d", &x, &y)
		*coordinates = append(*coordinates, coordinate{x, y})
	}
}

func maxBoundingCoordinate(coordinates []coordinate) coordinate {
	bounding := coordinate{0, 0}
	for _, coordinate := range coordinates {
		if coordinate.x > bounding.x {
			bounding.x = coordinate.x
		}
		if coordinate.y > bounding.y {
			bounding.y = coordinate.y
		}
	}
	return bounding
}
