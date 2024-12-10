// Advent of Code, 2024, Day 5
//
// https://adventofcode.com/2024/day/5
//
// Part 1 idea:
// The rule set given to us defines all possible pairs of pages. Knowing this,
// we can create a graph where the keys are the "from" page, and the values are
// the "to" page. Then we can just check if each adjacent pair is a rule
// violation.
package day05

import (
	"advent/util"
	"bufio"
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

func (e *Edge) Copy() *Edge {
	return NewEdge(e.From, e.To)
}

func PartOneAnswer(filepath string) (int, error) {
	return getValidOrderSum(filepath)
}

func PartTwoAnswer(filepath string) (int, error) {
	return 0, nil
}

func getValidOrderSum(filepath string) (int, error) {
	validOrderingSum := 0
	err := util.ProcessFile(filepath, func(s *bufio.Scanner) error {
		edges := getEdges(s)
		graph := getGraph(edges)
		for s.Scan() {
			line := s.Text()
			ordering := strings.Split(line, ",")
			if isValidOrdering(ordering, graph) {
				median, err := getValidOrderingMedian(ordering)
				if err != nil {
					return err
				}
				validOrderingSum += median
			}
		}
		return nil
	})
	return validOrderingSum, err
}

// getEdges takes a scanner and returns a list of all edges computed from the
// input file.
func getEdges(s *bufio.Scanner) []*Edge {
	edges := make([]*Edge, 0)
	for s.Scan() {
		line := s.Text()
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
func getGraph(edges []*Edge) Graph {
	graph := make(map[string]map[string]bool)
	for _, edge := range edges {
		if _, ok := graph[edge.From]; !ok {
			graph[edge.From] = make(map[string]bool, 0)
		}
		graph[edge.From][edge.To] = true
	}
	return graph
}

func isValidOrdering(ordering []string, graph Graph) bool {
	for i := 0; i < len(ordering)-1; i++ {
		if _, ok := graph[ordering[i+1]][ordering[i]]; ok {
			return false
		}
	}
	return true
}

func getValidOrderingMedian(ordering []string) (int, error) {
	median := ordering[len(ordering)/2]
	return strconv.Atoi(median)
}
