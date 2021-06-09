/*
Package trie implements a simple library for the Trie data structure
*/
package main

// TrieNode represents a single node in a Trie
type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

// NewTrieNode creates a new pre-defined trie node
// This method should only be called by trie library methods
// Users should not call this method directly
func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isWord:   false,
	}
}

// DFSCount implements a depth-first search count of words in a trie node
// This method should only be called by trie library methods
// Users should not call this method directly
func (node *TrieNode) DFSCount() int {
	count := 0
	if node.isWord {
		count = 1
	}
	for _, child := range node.children {
		count += child.DFSCount()
	}
	return count
}

// DFSList implements a depth-first search listing of words in a trie node
// This method should only be called by trie library methods
// Users should not call this method directly
func (node *TrieNode) DFSList(path []rune) []string {
	var words []string
	if node.isWord {
		words = append(words, string(path))
	}
	for char, child := range node.children {
		words = append(words, child.DFSList(append(path, char))...)
	}

	return words
}

// Trie is a simple struct that will only hold a root TrieNode
type Trie struct {
	root *TrieNode
}

// NewTrie creates a pre-defined empty Trie struct
func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

// Insert will accept a string and insert into a Trie struct
func (trie *Trie) Insert(word string) {
	currentNode := trie.root
	for _, char := range word {
		if currentNode.children[char] == nil {
			currentNode.children[char] = NewTrieNode()
		}
		currentNode = currentNode.children[char]
	}
	currentNode.isWord = true
}

// Contains parses a Trie and returns true if the trie contains the selected word
func (trie *Trie) Contains(word string) bool {
	currentNode := trie.root
	for _, char := range word {
		if _, ok := currentNode.children[char]; ok {
			currentNode = currentNode.children[char]
		} else {
			return false
		}
	}
	return currentNode.isWord
}

// Count parses a Trie and returns the number of words in the Trie
func (trie *Trie) Count() int {
	currentNode := trie.root
	return currentNode.DFSCount()
}

// List parses a Trie and returns a list of strings contained in the Trie
func (trie *Trie) List() []string {
	currentNode := trie.root
	return currentNode.DFSList(make([]rune, 0))
}

func (trie *Trie) Search(partial string) {

}
