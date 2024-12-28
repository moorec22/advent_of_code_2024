// Advent of Code, 2024, Day 22
//
// https://adventofcode.com/2024/day/22
//
// Part 1: I decided to go with the naive solution, calculating each number
// in each sequence with int math. In a future iteration, I will switch to
// bit math to try to shorten the process.
package day22

import (
	"advent/util"
	"bufio"
	"strconv"
)

// 2^24
const PruneNumber = 16777216

const FirstShiftUp = 6
const FirstShiftDown = 5
const SecondShiftUp = 11

const DailyNewSecrets = 2000

type Day22Solution struct {
	initialSecrets []int
}

func NewDay22Solution(filename string) (*Day22Solution, error) {
	initialSecrets := make([]int, 0)
	err := util.ProcessFile(filename, func(scanner *bufio.Scanner) error {
		for scanner.Scan() {
			if scanner.Text() == "" {
				continue
			}
			secret, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return err
			}
			initialSecrets = append(initialSecrets, secret)
		}
		return scanner.Err()
	})
	return &Day22Solution{initialSecrets}, err
}

func (s *Day22Solution) PartOneAnswer() (int, error) {
	newSecrets := make([]int, len(s.initialSecrets))
	for i, secret := range s.initialSecrets {
		newSecrets[i] = s.nthSecret(secret, DailyNewSecrets)
	}
	return util.SliceSum(newSecrets), nil
}

func (s *Day22Solution) PartTwoAnswer() (int, error) {
	return 0, nil
}

func (s *Day22Solution) nthSecret(secret, n int) int {
	for range n {
		secret = s.nextSecret(secret)
	}
	return secret
}

func (s *Day22Solution) nextSecret(secret int) int {
	secret = s.mixAndPruneNumber(secret, secret<<FirstShiftUp)
	secret = s.mixAndPruneNumber(secret, secret>>FirstShiftDown)
	secret = s.mixAndPruneNumber(secret, secret<<SecondShiftUp)

	return secret
}

// mixAndPruneNumber returns pruneNumber(mixNumber(n, m))
func (s *Day22Solution) mixAndPruneNumber(n, m int) int {
	return s.pruneNumber(s.mixNumber(n, m))
}

// mixNumber returns n ^ m
func (s *Day22Solution) mixNumber(n, m int) int {
	return n ^ m
}

// pruneNumber returns n % PruneNumber
func (s *Day22Solution) pruneNumber(n int) int {
	return n & (PruneNumber - 1)
}
