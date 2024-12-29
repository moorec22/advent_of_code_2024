// Advent of Code, 2024, Day 22
//
// https://adventofcode.com/2024/day/22
//
// Part 1: I decided to go with the naive solution, calculating each number
// in each sequence with int math. In a future iteration, I will switch to
// bit math to try to shorten the process.
//
// Part 2: If we take the first array of costs, we can create a sort of trie
// storing all our change sequences and associated bananas for efficient
// find and changing. We can then make a new trie with every array following,
// adding only the sequences that exist in that array too.
//
// It's a lot of effort, but I mostly want to try building the trie!
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

const SequenceLength = 4

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
	sequenceTrie := NewSequenceTrie(SequenceLength)
	for _, initialSecret := range s.initialSecrets {
		prices := s.getPrices(initialSecret, DailyNewSecrets)
		s.addNewPrices(sequenceTrie, prices)
	}
	return sequenceTrie.MaxBananas(), nil
}

// getPrices returns n prices in an array. initialSecret is the first secret in
// the prices for the day.
func (s *Day22Solution) getPrices(initialSecret, n int) []int {
	prices := make([]int, n)
	secret := initialSecret
	prices[0] = secret % 10
	for i := 1; i < n; i++ {
		secret = s.nextSecret(secret)
		prices[i] = secret % 10
	}
	return prices
}

// nthSecret returns the nth secret in a sequence
func (s *Day22Solution) nthSecret(secret, n int) int {
	for range n {
		secret = s.nextSecret(secret)
	}
	return secret
}

// nextSecret takes a secret and returns the next secret
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

// addNewPrices will add the sequences found in prices to trie, using
// SequenceLength to determine the length of sequences. If the sequence does
// not yet exist, the banana value in prices is used. Otherwise, the banana
// value in prices is added on.
func (s *Day22Solution) addNewPrices(trie *SequenceTrie, prices []int) {
	sequence := util.NewArrayQueue[int]()
	newTrie := NewSequenceTrie(SequenceLength)
	for i := 1; i < len(prices); i++ {
		sequence.Insert(prices[i] - prices[i-1])
		if sequence.Size() == SequenceLength {
			newTrie.Insert(sequence.ToArray(), prices[i])
			sequence.Remove()
		}
	}
	trie.MergeInto(newTrie)
}
