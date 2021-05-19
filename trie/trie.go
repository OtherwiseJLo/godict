package trie

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

type NodeDequeue struct {
	nodes []*TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isWord:   false,
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

func dfs(node *TrieNode, path []rune) int {
	count := 0
	if node.isWord {
		count = 1
	}
	for char, child := range node.children {
		count += dfs(child, append(path, char))
	}
	return count
}

func (trie *Trie) Count() int {
	currentNode := trie.root
	return dfs(currentNode, make([]rune, 1))
}
