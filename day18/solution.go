// Advent of Code, 2024, Day 18
//
// https://adventofcode.com/2024/day/18
//
// Part 1: I decided to finally implement a priority queue in util. From there,
// I used djikstra's algorithm to find the shortest path from the start to the.
// end. I used a matrix to represent the memory space and a matrix of CellInfo
// to represent the shortest path to each cell.
package day18

import (
	"advent/util"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

const MemoryHeight = 71
const MemoryWidth = 71

const ByteCount = 1024

type CellInfo struct {
	pos          util.Vector
	shortestPath int
	sym          rune
}

func (c CellInfo) Compare(other CellInfo) int {
	return c.shortestPath - other.shortestPath
}

type Day18Solution struct {
	memorySpace  util.Matrix[rune]
	fallingBytes []*util.Vector
}

func NewDay18Solution(filename string) (*Day18Solution, error) {
	memorySpace := util.NewMatrix[rune]()
	for i := 0; i < MemoryHeight; i++ {
		memorySpace = append(memorySpace, make([]rune, MemoryWidth))
		for j := 0; j < MemoryWidth; j++ {
			memorySpace[i][j] = '.'
		}
	}
	fallingBytes := make([]*util.Vector, 0)
	err := util.ProcessFile(filename, func(scanner *bufio.Scanner) error {
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Split(line, ",")
			y, err := strconv.Atoi(parts[0])
			if err != nil {
				return err
			}
			x, err := strconv.Atoi(parts[1])
			if err != nil {
				return err
			}
			fallingBytes = append(fallingBytes, util.NewVector(x, y))
		}
		return scanner.Err()
	})
	return &Day18Solution{memorySpace, fallingBytes}, err
}

func (s *Day18Solution) PartOneAnswer() (int, error) {
	s.simulateXBytes(s.memorySpace, ByteCount, s.fallingBytes)
	start := util.NewVector(0, 0)
	end := util.NewVector(MemoryWidth-1, MemoryHeight-1)
	shortestPath := s.findShortestPath(s.memorySpace, start, end)
	if shortestPath == -1 {
		return -1, fmt.Errorf("no path found")
	}
	return shortestPath, nil
}

func (s *Day18Solution) PartTwoAnswer() (int, error) {
	return 0, nil
}

// simulateXBytes simulates the first x bytes from bytes falling into thememry
// space.
func (s *Day18Solution) simulateXBytes(memorySpace util.Matrix[rune], x int, bytes []*util.Vector) {
	for i := 0; i < x; i++ {
		position := bytes[i]
		memorySpace.Set(position, '#')
	}
}

// findShortestPath finds the shortest path from the start to the end in the
// memory space.
func (s *Day18Solution) findShortestPath(memorySpace util.Matrix[rune], start, end *util.Vector) int {
	memorySpaceMap := s.getMemorySpaceMap(memorySpace)
	pq := util.NewArrayPriorityQueue[CellInfo]()
	visited := make(map[util.Vector]bool)
	pq.Insert(CellInfo{pos: *start, shortestPath: 0, sym: '.'})
	for !pq.IsEmpty() {
		current := pq.Remove()
		if current.pos.Equals(end) {
			return current.shortestPath
		}
		if visited[current.pos] {
			continue
		}
		visited[current.pos] = true
		neighbors := s.getValidNeighbors(memorySpaceMap, &current.pos)
		for _, neighbor := range neighbors {
			neighborInfo := memorySpaceMap.Get(neighbor)
			shortestPath := current.shortestPath + 1
			if !visited[*neighbor] && (neighborInfo.shortestPath == -1 || neighborInfo.shortestPath > shortestPath) {
				neighborInfo.shortestPath = shortestPath
				memorySpaceMap.Set(neighbor, neighborInfo)
				pq.Insert(neighborInfo)
			}
		}
	}
	return -1
}

// getValidNeighbors returns the valid neighbors of a position in the memory
// space. A valid neighbor is a position that is within the bounds of the memory
// space and is a '.'.
func (s *Day18Solution) getValidNeighbors(memorySpace util.Matrix[CellInfo], position *util.Vector) []*util.Vector {
	neighbors := make([]*util.Vector, 0)
	for _, direction := range util.SimpleDirections {
		neighbor := position.Add(direction)
		if memorySpace.PosInBounds(neighbor) && memorySpace.Get(neighbor).sym == '.' {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

func (s *Day18Solution) getMemorySpaceMap(memorySpace util.Matrix[rune]) util.Matrix[CellInfo] {
	memorySpaceMap := util.NewMatrix[CellInfo]()
	for i := 0; i < MemoryHeight; i++ {
		memorySpaceMap = append(memorySpaceMap, make([]CellInfo, MemoryWidth))
		for j := 0; j < MemoryWidth; j++ {
			position := util.NewVector(i, j)
			if memorySpace.Get(position) == '.' {
				memorySpaceMap.Set(position, CellInfo{pos: *position, shortestPath: -1, sym: '.'})
			}
		}
	}
	return memorySpaceMap
}
