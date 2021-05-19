package main

type Trie struct {
	root TrieNode
}

type TrieNode struct {
	children map[rune]TrieNode
	isWord   bool
}

func (trie Trie) Insert(word string) {
	currentNode := trie.root
	for idx, char := range word {
		if _, ok := currentNode.children[char]; ok {
			currentNode = currentNode.children[char]
		} else {
			currentNode.children[char] = TrieNode{
				children: make(map[rune]TrieNode),
				isWord:   idx == len(word)+1,
			}
			currentNode = currentNode.children[char]
		}
	}
}

func (trie Trie) Contains(word string) bool {
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
