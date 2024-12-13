// Advent of Code, 2024, Day 5
//
// https://adventofcode.com/2024/day/5
//
// Part 1 idea:
// The rule set given to us defines all possible pairs of pages. Knowing this,
// we can create a graph where the keys are the "from" page, and the values are
// the "to" page. Then we can just check if each adjacent pair is a rule
// violation.
//
// Part 2 idea:
// Because they give us all possible pairs of pages (and given the prompt), we
// can sort a list by comparing two pages based on their edges in the graph.
// For two pages, if p -> q is in the graph, then p is "smaller". Otherwise,
// q is "smaller".
package day05

import (
	"advent/util"
	"bufio"
	"slices"
	"strconv"
	"strings"
)

type Graph map[string]map[string]bool

type Edge struct {
	From string
	To   string
}

func NewEdge(from, to string) *Edge {
	return &Edge{From: from, To: to}
}

type Day05Solution struct {
	filepath string
}

func NewDay05Solution(filepath string) (*Day05Solution, error) {
	return &Day05Solution{filepath: filepath}, nil
}

func (s *Day05Solution) PartOneAnswer() (int, error) {
	validOrderingsSum := 0
	err := util.ProcessFile(s.filepath, func(scanner *bufio.Scanner) error {
		edges := s.getEdges(scanner)
		graph := s.getGraph(edges)
		validOrderings, err := s.getOrderings(scanner, graph, true)
		if err != nil {
			return err
		}
		validOrderingsSum, err = s.getSumOfMedians(validOrderings)
		return err
	})
	if err != nil {
		return 0, err
	}
	return validOrderingsSum, err
}

func (s *Day05Solution) PartTwoAnswer() (int, error) {
	reorderedSum := 0
	err := util.ProcessFile(s.filepath, func(scanner *bufio.Scanner) error {
		edges := s.getEdges(scanner)
		graph := s.getGraph(edges)
		invalidOrderings, err := s.getOrderings(scanner, graph, false)
		if err != nil {
			return err
		}
		reorderedOrderings := make([][]string, 0)
		for _, invalidOrdering := range invalidOrderings {
			reorderedOrderings = append(reorderedOrderings, s.getValidOrdering(invalidOrdering, graph))
		}
		reorderedSum, err = s.getSumOfMedians(reorderedOrderings)
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		return 0, err
	}
	return reorderedSum, err
}

// getOrderingsTwo takes a scanner and returns a list of all orderings computed
// from the input file. If uses the graph to determine validity, and returns all
// valid orderings if valid is true, and all invalid orderings if valid is false.
func (s *Day05Solution) getOrderings(scanner *bufio.Scanner, graph Graph, valid bool) ([][]string, error) {
	orderings := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		ordering := strings.Split(line, ",")
		if valid == s.isValidOrdering(ordering, graph) {
			orderings = append(orderings, ordering)
		}
	}
	return orderings, nil
}

func (s *Day05Solution) getSumOfMedians(lists [][]string) (int, error) {
	medianSum := 0
	for _, list := range lists {
		median, err := s.getValidOrderingMedian(list)
		if err != nil {
			return 0, err
		}
		medianSum += median
	}
	return medianSum, nil
}

// getEdges takes a scanner and returns a list of all edges computed from the
// input file.
func (s *Day05Solution) getEdges(scanner *bufio.Scanner) []*Edge {
	edges := make([]*Edge, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, "|")
		edges = append(edges, NewEdge(parts[0], parts[1]))
	}
	return edges
}

// getGraph takes a list of edges, and returns a graph representation, where
// keys are the To field of the edge, and the values are the From field of the
// edge.
func (s *Day05Solution) getGraph(edges []*Edge) Graph {
	graph := make(map[string]map[string]bool)
	for _, edge := range edges {
		if _, ok := graph[edge.From]; !ok {
			graph[edge.From] = make(map[string]bool, 0)
		}
		graph[edge.From][edge.To] = true
	}
	return graph
}

func (s *Day05Solution) isValidOrdering(ordering []string, graph Graph) bool {
	for i := 0; i < len(ordering)-1; i++ {
		if _, ok := graph[ordering[i+1]][ordering[i]]; ok {
			return false
		}
	}
	return true
}

func (s *Day05Solution) getValidOrderingMedian(ordering []string) (int, error) {
	median := ordering[len(ordering)/2]
	return strconv.Atoi(median)
}

func (s *Day05Solution) getValidOrdering(ordering []string, graph Graph) []string {
	slices.SortFunc(ordering, func(i, j string) int {
		if _, ok := graph[j][i]; ok {
			return 1
		} else {
			return -1
		}
	})
	return ordering
}
