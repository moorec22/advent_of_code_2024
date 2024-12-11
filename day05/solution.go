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

func PartOneAnswer(filepath string) (int, error) {
	validOrderingsSum := 0
	err := util.ProcessFile(filepath, func(s *bufio.Scanner) error {
		edges := getEdges(s)
		graph := getGraph(edges)
		validOrderings, err := getOrderings(s, graph, true)
		if err != nil {
			return err
		}
		validOrderingsSum, err = getSumOfMedians(validOrderings)
		return err
	})
	if err != nil {
		return 0, err
	}
	return validOrderingsSum, err
}

func PartTwoAnswer(filepath string) (int, error) {
	reorderedSum := 0
	err := util.ProcessFile(filepath, func(s *bufio.Scanner) error {
		edges := getEdges(s)
		graph := getGraph(edges)
		invalidOrderings, err := getOrderings(s, graph, false)
		if err != nil {
			return err
		}
		reorderedOrderings := make([][]string, 0)
		for _, invalidOrdering := range invalidOrderings {
			reorderedOrderings = append(reorderedOrderings, getValidOrdering(invalidOrdering, graph))
		}
		reorderedSum, err = getSumOfMedians(reorderedOrderings)
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
func getOrderings(s *bufio.Scanner, graph Graph, valid bool) ([][]string, error) {
	orderings := make([][]string, 0)
	for s.Scan() {
		line := s.Text()
		ordering := strings.Split(line, ",")
		if valid == isValidOrdering(ordering, graph) {
			orderings = append(orderings, ordering)
		}
	}
	return orderings, nil
}

func getSumOfMedians(lists [][]string) (int, error) {
	medianSum := 0
	for _, list := range lists {
		median, err := getValidOrderingMedian(list)
		if err != nil {
			return 0, err
		}
		medianSum += median
	}
	return medianSum, nil
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

func getValidOrdering(ordering []string, graph Graph) []string {
	slices.SortFunc(ordering, func(i, j string) int {
		if _, ok := graph[j][i]; ok {
			return 1
		} else {
			return -1
		}
	})
	return ordering
}
