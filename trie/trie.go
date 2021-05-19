package trie

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			isWord:   false,
		},
	}
}

func (trie *Trie) Insert(word string) {
	currentNode := trie.root
	for _, char := range word {
		if currentNode.children[char] == nil {
			currentNode.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
				isWord:   false,
			}
		}
		currentNode = currentNode.children[char]
	}
	currentNode.isWord = true
}

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
