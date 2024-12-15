// Advent of Code, 2024, Day 8
//
// https://adventofcode.com/2024/day/8
//
// Part 1: I decided to add manhattan distances to my utilities package. After
// that, it was pretty easy to calculate the antinodes for the fixed antennas.
// Just add the manhattan distance to either side of both antennas.
//
// Part 2: To find resonant antinodes, we have to make sure we have the unit
// manhattan distance. We can then add the unit manhattan distance to each
// antenna's Vector until we reach the edge of the city map.
package day08

import (
	"advent/util"
)

const Empty = '.'

type Antenna struct {
	Vector util.Vector
	symbol rune
}

type Day08Solution struct {
	cityMap  util.Matrix[rune]
	antennas map[rune][]Antenna
}

func NewDay08Solution(filepath string) (*Day08Solution, error) {
	cityMap, err := util.ParseMatrix(filepath, func(r rune) rune {
		return r
	})
	if err != nil {
		return nil, err
	}
	antennas := getAntennas(cityMap)
	return &Day08Solution{cityMap, antennas}, nil
}

func (s *Day08Solution) PartOneAnswer() (int, error) {
	antinodes := s.getAntinodesVectors(s.cityMap, s.antennas, s.getFixedAntinodes)
	return len(antinodes), nil
}

func (s *Day08Solution) PartTwoAnswer() (int, error) {
	antinodes := s.getAntinodesVectors(s.cityMap, s.antennas, s.getResonantAntinodes)
	return len(antinodes), nil
}

// getAntennas returns a map of antennas by their symbol.
func getAntennas(cityMap util.Matrix[rune]) map[rune][]Antenna {
	antennas := make(map[rune][]Antenna)
	for y, row := range cityMap {
		for x, symbol := range row {
			if symbol != Empty {
				antennas[symbol] = append(antennas[symbol], Antenna{util.NewVector(x, y), symbol})
			}
		}
	}
	return antennas
}

// getAntinodesVectors returns all antinodes defined by the list of antennas. It uses getAntinodesFromAntennas
// to calculate the antinodes for each pair of antennas.
func (s *Day08Solution) getAntinodesVectors(cityMap util.Matrix[rune], antennasBySymbol map[rune][]Antenna,
	getAntinodesFromAntennas func(Antenna, Antenna, util.Matrix[rune]) []util.Vector) map[util.Vector]bool {
	antinodes := make(map[util.Vector]bool)
	for _, antennas := range antennasBySymbol {
		for i, antenna1 := range antennas {
			for j := i + 1; j < len(antennas); j++ {
				antenna2 := antennas[j]
				antinodesFromAntennas := getAntinodesFromAntennas(antenna1, antenna2, cityMap)
				for _, antinode := range antinodesFromAntennas {
					if cityMap.PosInBounds(antinode) {
						antinodes[antinode] = true
					}
				}
			}
		}
	}
	return antinodes
}

// getFixedAntinodes returns all fixed antinodes defined by two antennas.
// As defined by the problem, fixed antinodes can be in two Vectors: the two
// Vectors that are exactly twice as far from one antenna as the other.
func (s *Day08Solution) getFixedAntinodes(one, two Antenna, cityMap util.Matrix[rune]) []util.Vector {
	fixedAntinodes := make([]util.Vector, 0)
	manhattanDistance := two.Vector.GetManhattanDistance(one.Vector)
	negativeManhattanDistance := manhattanDistance.Negate()
	oneAntinode := one.Vector.AddManhattanDistance(negativeManhattanDistance)
	twoAntinode := two.Vector.AddManhattanDistance(manhattanDistance)
	if s.cityMap.PosInBounds(oneAntinode) {
		fixedAntinodes = append(fixedAntinodes, oneAntinode)
	}
	if s.cityMap.PosInBounds(twoAntinode) {
		fixedAntinodes = append(fixedAntinodes, twoAntinode)
	}
	return fixedAntinodes
}

// getResonantAntinodes returns all resonant antinodes defined by two antennas.
// As defined by the problem, resonant antiodes are any antinodes exactly in
// line with the two antennas.
func (s *Day08Solution) getResonantAntinodes(one, two Antenna, cityMap util.Matrix[rune]) []util.Vector {
	resonantAntinodes := make([]util.Vector, 0)
	manhattanDistance := two.Vector.GetManhattanDistance(one.Vector)
	unitManhattanDistance := manhattanDistance.Unit()
	negativeUnitManhattanDistance := unitManhattanDistance.Negate()
	resonantFromOne := one.Vector
	for s.cityMap.PosInBounds(resonantFromOne) {
		resonantAntinodes = append(resonantAntinodes, resonantFromOne)
		resonantFromOne = resonantFromOne.AddManhattanDistance(negativeUnitManhattanDistance)
	}
	resonantFromTwo := two.Vector
	for s.cityMap.PosInBounds(resonantFromTwo) {
		resonantAntinodes = append(resonantAntinodes, resonantFromTwo)
		resonantFromTwo = resonantFromTwo.AddManhattanDistance(unitManhattanDistance)
	}
	return resonantAntinodes
}
