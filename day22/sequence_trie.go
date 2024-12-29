package day22

import "fmt"

// SequenceNode is a node of a sequence.
type SequenceNode struct {
	bananas  int
	children map[int]*SequenceNode
}

func NewSequenceNode() *SequenceNode {
	return &SequenceNode{-1, make(map[int]*SequenceNode)}
}

// A trie storing a sequence and cost.
type SequenceTrie struct {
	root   *SequenceNode
	length int
}

func NewSequenceTrie(length int) *SequenceTrie {
	return &SequenceTrie{NewSequenceNode(), length}
}

// Insert inserts the sequence into the trie, if the sequence is not yet in the
// trie.
func (t *SequenceTrie) Insert(sequence []int, bananas int) error {
	if len(sequence) != t.length {
		return fmt.Errorf("sequence is not of length %d", t.length)
	}
	node := t.root
	for _, digit := range sequence {
		if _, ok := node.children[digit]; !ok {
			node.children[digit] = NewSequenceNode()
		}
		node = node.children[digit]
	}
	if node.bananas == -1 {
		node.bananas = bananas
	}
	return nil
}

func (t *SequenceTrie) Bananas(sequence []int) (int, error) {
	leaf := t.leaf(sequence)
	if leaf == nil {
		return 0, fmt.Errorf("sequence is not in trie: %v", sequence)
	} else {
		return leaf.bananas, nil
	}
}

// returns the maximum bananas count in the trie.
func (t *SequenceTrie) MaxBananas() int {
	return t.maxBananasHelper(t.root)
}

func (t *SequenceTrie) maxBananasHelper(node *SequenceNode) int {
	if len(node.children) == 0 {
		return node.bananas
	} else {
		maxBananas := -1
		for _, child := range node.children {
			bananas := t.maxBananasHelper(child)
			if maxBananas == -1 || maxBananas < bananas {
				maxBananas = bananas
			}
		}
		return maxBananas
	}
}

// merges other into this trie, combining existing banana counts.
func (t *SequenceTrie) MergeInto(other *SequenceTrie) error {
	if t.length != other.length {
		return fmt.Errorf("mismatched lengths for merging tries")
	}
	t.mergeIntoHelper(t.root, other.root)
	return nil
}

func (t *SequenceTrie) mergeIntoHelper(ourNode, otherNode *SequenceNode) {
	if len(otherNode.children) == 0 {
		ourNode.bananas += otherNode.bananas
		return
	}
	for digit, child := range otherNode.children {
		if _, ok := ourNode.children[digit]; !ok {
			ourNode.children[digit] = child
		} else {
			t.mergeIntoHelper(ourNode.children[digit], otherNode.children[digit])
		}
	}
}

func (t *SequenceTrie) leaf(sequence []int) *SequenceNode {
	node := t.root
	for _, digit := range sequence {
		if _, ok := node.children[digit]; !ok {
			return nil
		}
		node = node.children[digit]
	}
	return node
}
