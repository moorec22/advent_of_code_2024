// Advent of Code, 2024, Day 12
//
// https://adventofcode.com/2024/day/12
//
// Part 1: I decided to store the garden in a 2D matrix of GardenSquare
// structs, which contain a rune for the plant and a boolean for whether
// the square has been visited. Then, for part 1, we just do a pass over
// the entire matrix, doing a depth-first search from each plant to determine
// the permiter and area of the garden region. I decided not to create
// the garden square structure in the initializer, since it's a temporary
// representation of the garden (overkill I know).
//
// Part 2: For part 2, we need to return the number of sides instead of the
// permiter. There must be exactly as many sides as there are corners for a
// closed polygon. So we can just find the number of corners instead.
package day12

import (
	"advent/util"
	"fmt"
)

type GardenSquare struct {
	Plant   rune
	visited bool
}

type Day12Solution struct {
	gardenMap util.Matrix[rune]
}

func NewDay12Solution(filepath string) (*Day12Solution, error) {
	gardenMap, err := util.ParseMatrixFromFile(filepath, func(r rune) rune {
		return r
	})
	return &Day12Solution{gardenMap}, err
}

func (s *Day12Solution) PartOneAnswer() (int, error) {
	gardenSquareMap := getGardenSquareMap(s.gardenMap)
	return s.getFencingPrice(gardenSquareMap), nil
}

func (s *Day12Solution) PartTwoAnswer() (int, error) {
	gardenSquareMap := getGardenSquareMap(s.gardenMap)
	return s.getFencingPriceWithDiscount(gardenSquareMap), nil
}

func (s *Day12Solution) getFencingPrice(gardenSquareMap util.Matrix[GardenSquare]) int {
	price := 0
	for i := range gardenSquareMap {
		for j := range gardenSquareMap[i] {
			if !gardenSquareMap[i][j].visited {
				area, perimiter := s.getFencingAreaAndPerimiter(gardenSquareMap, util.NewVector(i, j))
				price += area * perimiter
			}
		}
	}
	return price
}

func (s *Day12Solution) getFencingPriceWithDiscount(gardenSquareMap util.Matrix[GardenSquare]) int {
	price := 0
	for i := range gardenSquareMap {
		for j := range gardenSquareMap[i] {
			if !gardenSquareMap[i][j].visited {
				area, corners := s.getFencingAreaAndCorners(gardenSquareMap, util.NewVector(i, j))
				fmt.Println(string(gardenSquareMap[i][j].Plant), area, corners)
				price += area * corners
			}
		}
	}
	return price
}

// getFencingAreaAndPermiter returns the area and perimiter of the remaining
// part of the region starting at p that is unvisited. As specified in the
// problem, the permiter is the number of sides of the region that are adjacent
// to the edge of the garden or different garden plots, and the area is the
// total number of squares in the plot.
func (s *Day12Solution) getFencingAreaAndPerimiter(gardenMap util.Matrix[GardenSquare], p *util.Vector) (int, int) {
	currentSquare := gardenMap.Get(p)
	if currentSquare.visited {
		return 0, 0
	}
	gardenMap[p.X][p.Y].visited = true
	area := 1
	perimiter := 0
	for _, d := range util.SimpleDirections {
		newPos := p.Add(d)
		directionInRegion := gardenMap.PosInBounds(newPos) && gardenMap.Get(newPos).Plant == currentSquare.Plant
		if directionInRegion {
			newArea, newPerimiter := s.getFencingAreaAndPerimiter(gardenMap, newPos)
			area += newArea
			perimiter += newPerimiter
		} else {
			perimiter++
		}
	}
	return area, perimiter
}

// getFencingAreaAndCorners returns the area and number of corners of the
// remaining part of the region starting at p that is unvisited.
func (s *Day12Solution) getFencingAreaAndCorners(gardenMap util.Matrix[GardenSquare], p *util.Vector) (int, int) {
	currentSquare := gardenMap.Get(p)
	if currentSquare.visited {
		return 0, 0
	}
	gardenMap[p.X][p.Y].visited = true
	area := 1
	corners := 0
	for _, d := range util.SimpleDirections {
		newPos := p.Add(d)
		directionInRegion := gardenMap.PosInBounds(newPos) && gardenMap.Get(newPos).Plant == currentSquare.Plant
		if directionInRegion {
			newArea, newCorners := s.getFencingAreaAndCorners(gardenMap, newPos)
			area += newArea
			corners += newCorners
		}
	}

	// Check for corners
	directionsInRegion := make([]bool, len(util.AllDirections))
	for i, d := range util.AllDirections {
		newPos := p.Add(d)
		directionsInRegion[i] = gardenMap.PosInBounds(newPos) && gardenMap.Get(newPos).Plant == currentSquare.Plant
	}
	for i := 0; i < len(directionsInRegion); i += 2 {
		// convex corner
		firstDirection := directionsInRegion[i]
		secondDirection := directionsInRegion[(i+2)%len(directionsInRegion)]
		diagonalDirection := directionsInRegion[(i+1)%len(directionsInRegion)]
		// concave corner
		if firstDirection && secondDirection && !diagonalDirection {
			corners++
		}
		// convex corner
		if !firstDirection && !secondDirection {
			corners++
		}
	}
	return area, corners
}

// getGardenSquareMap returns a matrix of GardenSquare structs from a matrix of
// runes.
func getGardenSquareMap(gardenMap util.Matrix[rune]) util.Matrix[GardenSquare] {
	squareMap := make(util.Matrix[GardenSquare], len(gardenMap))
	for i := range gardenMap {
		squareMap[i] = make([]GardenSquare, len(gardenMap[i]))
		for j := range gardenMap[i] {
			squareMap[i][j] = GardenSquare{gardenMap[i][j], false}
		}
	}
	return squareMap
}
