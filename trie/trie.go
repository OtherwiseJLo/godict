package trie

type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isWord:   false,
	}
}

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

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

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

func (trie *Trie) Count() int {
	currentNode := trie.root
	return currentNode.DFSCount()
}

func (trie *Trie) List() []string {
	currentNode := trie.root
	return currentNode.DFSList(make([]rune, 0))
}
