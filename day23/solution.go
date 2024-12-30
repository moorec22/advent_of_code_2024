// Advent of Code, 2024, Day 23
//
// https://adventofcode.com/2024/day/23
//
// Part 1: A triplet of interconnected computers by definition are a cycle
// of three computers. If we search each node for its cycles of three, we
// can find all triplets. We can check for uniqueness by storing our cycles
// alphabetically in a single string.
//
// Part 2: Can we do a little induction programming here? What if we have all
// dense graphs of size X. Then take a node n. If a dense graph is a subset
// of n.connections, then {n, subset...} is a dense graph of size X plus
// 1. Then, we can rewrite part 1 in this new way.
package day23

import (
	"advent/util"
	"bufio"
	"fmt"
	"slices"
	"strings"
)

type Graph map[string]map[string]bool

type Day23Solution struct {
	lanGraph Graph
}

func NewDay23Solution(filename string) (*Day23Solution, error) {
	lanGraph := make(Graph)
	err := util.ProcessFile(filename, func(scanner *bufio.Scanner) error {
		for scanner.Scan() {
			connection := strings.Split(scanner.Text(), "-")
			if len(connection) != 2 {
				return fmt.Errorf("incorrectly formatted line in file: %s", scanner.Text())
			}
			addGraphConnection(lanGraph, connection[0], connection[1])
			addGraphConnection(lanGraph, connection[1], connection[0])
		}
		return scanner.Err()
	})
	return &Day23Solution{lanGraph}, err
}

func (s *Day23Solution) PartOneAnswer() (int, error) {
	denseSets := s.getDenseSetsOfSize(s.lanGraph, 3)
	denseSets = s.filterForPrefix(denseSets, "t")
	return len(denseSets), nil
}

func (s *Day23Solution) PartTwoAnswer() (int, error) {
	largestDenseSets := s.getLargestDenseNodeSets(s.lanGraph)
	if len(largestDenseSets) != 1 {
		return 0, fmt.Errorf("there should be exactly one largest dense set, not %d", len(largestDenseSets))
	}
	fmt.Println(s.getPassword(largestDenseSets[0]))
	return 0, nil
}

// getDenseSetsOfSize returns all fully dense sets found in graph of size size.
func (s *Day23Solution) getDenseSetsOfSize(graph Graph, size int) []map[string]bool {
	denseSets := s.getFirstDenseSets(graph)
	for range size - 1 {
		denseSets = s.getDenseSets(graph, denseSets)
	}
	return denseSets
}

// filterForPrefix returns a filtered list of all sets in sets that have at
// least one string starting with prefix.
func (s *Day23Solution) filterForPrefix(sets []map[string]bool, prefix string) []map[string]bool {
	newSets := make([]map[string]bool, 0)
	for _, set := range sets {
		if s.containsPrefix(set, prefix) {
			newSets = append(newSets, set)
		}
	}
	return newSets
}

// getLargestDenseNodeSets returns a list of all node sets that are the largest
// dense node sets found in graph.
func (s *Day23Solution) getLargestDenseNodeSets(graph Graph) []map[string]bool {
	denseSets := s.getFirstDenseSets(graph)
	nextDenseSet := s.getDenseSets(graph, denseSets)
	for len(nextDenseSet) > 0 {
		denseSets = nextDenseSet
		nextDenseSet = s.getDenseSets(graph, nextDenseSet)
	}
	return denseSets
}

// getDenseSets takes a graph and searches for fully dense set within
// graph. A fully dense set is a set of nodes where all nodes are connected in
// graph. oldDenseSets is full of previous dense sets found of size X. It is
// assumed that oldDenseSets is the complete set of dense sets found of size X.
// Given that, all dense sets of size X + 1 are returned. If no dense sets have
// previously been found, set oldDenseSets to be the set of all nodes (which
// are also all dense sets of size 1)
func (s *Day23Solution) getDenseSets(graph Graph, oldDenseSets []map[string]bool) []map[string]bool {
	denseSets := make([]map[string]bool, 0)
	for node, connections := range graph {
		for _, oldDenseSet := range oldDenseSets {
			if _, ok := oldDenseSet[node]; !ok && s.isSubsetOf(connections, oldDenseSet) {
				// node is not in the oldDenseSet, and is connected to all nodes in oldDenseSet
				denseSet := util.CopyMap(oldDenseSet)
				denseSet[node] = true
				// this takes forever, what's a faster way to think about this?
				if !s.nodeSetContains(denseSets, denseSet) {
					denseSets = append(denseSets, denseSet)
				}
			}
		}
	}
	return denseSets
}

// nodeSetContains returns true if and only if nodeSets contains nodeSet.
func (s *Day23Solution) nodeSetContains(nodeSets []map[string]bool, nodeSet map[string]bool) bool {
	for _, ns := range nodeSets {
		if len(ns) == len(nodeSet) && s.isSubsetOf(nodeSet, ns) {
			return true
		}
	}
	return false
}

// isSubsetOf returns true if and only if setTwo is a subset of setOne
func (s *Day23Solution) isSubsetOf(setOne, setTwo map[string]bool) bool {
	for el := range setTwo {
		if _, ok := setOne[el]; !ok {
			return false
		}
	}
	return true
}

// containsPrefix returns true if and only if set contains at least one string
// with the given prefix
func (s *Day23Solution) containsPrefix(set map[string]bool, prefix string) bool {
	for s := range set {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

// getFirstDenseSets takes a graph, and returns the dense sets of size one.
// This is just a list of every node in its own set.
func (s *Day23Solution) getFirstDenseSets(graph Graph) []map[string]bool {
	denseSets := make([]map[string]bool, 0)
	for node := range graph {
		denseSets = append(denseSets, map[string]bool{node: true})
	}
	return denseSets
}

// getPassword returns every node in nodes, sorted alphabetically, and
// separated by commas.
func (s *Day23Solution) getPassword(nodes map[string]bool) string {
	nodeList := make([]string, 0)
	for node := range nodes {
		nodeList = append(nodeList, node)
	}
	slices.Sort(nodeList)
	return strings.Join(nodeList, ",")
}

// addGraphConnection makes an asymmetric connection from left to right in
// graph.
func addGraphConnection(graph Graph, left, right string) {
	if _, ok := graph[left]; !ok {
		graph[left] = make(map[string]bool)
	}
	graph[left][right] = true
}
