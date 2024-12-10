// Advent of Code, 2024, Day 5
//
// https://adventofcode.com/2024/day/5
//
// I decided to do part 1 by preprocessing the rules into a graph, and then
// getting a list of all rules that were once indirect using the transitive
// property. i.e., if a -> b -> c, then a -> c, where a, b, and c are rules.
// This processing can take O(v^2) time, where v is the number of pages. But
// afterwards the verficiation of each ordering takes O(n) time, where n is the
// number of pages in the ordering.
package day05

import (
	"advent/util"
	"bufio"
	"fmt"
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
		graph := getDenseGraph(edges)
		fmt.Println(graph)
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

// getDenseGraph takes a list of edges and returns a graph where edges are
// rules, indirect or direct, between nodes. For example, an edge may represent
// rule a -> b, and another edge may represent rule a-> b -> c.
func getDenseGraph(edges []*Edge) Graph {
	toProcess := make([]*Edge, 0)
	denseGraph := getGraph(edges)
	for _, edge := range edges {
		toProcess = append(toProcess, edge.Copy())
	}
	for len(toProcess) > 0 {
		edge := toProcess[0]
		toProcess = toProcess[1:]
		firstLevel := denseGraph[edge.From]
		secondLevel := denseGraph[edge.To]
		for node := range secondLevel {
			if _, ok := firstLevel[node]; !ok {
				// adding the second level indirect connection as a direct connection
				firstLevel[node] = true
				// setting new node to process
				toProcess = append(toProcess, NewEdge(edge.From, node))
			}
		}
		denseGraph[edge.From] = firstLevel
	}
	return denseGraph
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
			fmt.Println(ordering, ordering[i], ordering[i+1], graph[ordering[i+1]])
			return false
		}
	}
	return true
}

func getValidOrderingMedian(ordering []string) (int, error) {
	median := ordering[len(ordering)/2]
	return strconv.Atoi(median)
}
