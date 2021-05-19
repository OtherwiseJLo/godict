package trie

import "testing"

var words = []string{"build", "builds", "graph", "graphite", "graphing", "gran", "granular", "graphic", "a", "an", "and", "andes", "intellect", "intel", "int", "in", "i"}

func contains(words []string, partial string) bool {
	for _, word := range words {
		if word == partial {
			return true
		}
	}
	return false
}

// var bow BagOfWords = NewBagOfWords("build", "builds", "graph", "graphite", "graphing", "gran", "granular", "graphic", "a", "an", "and", "andes", "intellect", "intel", "int", "in", "i")

func TestInsert(t *testing.T) {
	trie := NewTrie()

	for testIdx, word := range words {
		t.Logf("%d2 TestInsert: %s", testIdx, word)
		trie.Insert(word)
		node := trie.root
		for idx, char := range word {
            t.Log(char)
			t.Log(node.children)
			if _, ok := node.children[char]; ok {
				node = node.children[char]
			} else {
				t.Errorf("could not find char %c (%d) in %s", idx, char, word)
			}
		}
	}
}

func TestContains(t *testing.T) {
	trie := NewTrie()

	for testIdx, word := range words {
		t.Logf("%d TestContains: %s", testIdx, word)
		trie.Insert(word)

		var path []rune
		for idx, char := range word {
			path = append(path, char)
			partialWord := string(path)
			if trie.Contains(partialWord) {
				if idx != len(word)-1 {
					if !contains(words, partialWord) {
						t.Errorf("%s: %s should not be in trie", word, partialWord)
					}
				}
			}
		}
		if trie.Contains(word) {
			t.Logf("%s found in trie", word)
		} else {
			t.Errorf("could not find %s in trie", word)
		}
	}
}
