package trie

import "testing"

type BagOfWords struct {
	words map[string]int
}

func NewBagOfWords(words ...string) BagOfWords {
	bow := make(map[string]int)
	for idx, word := range words {
		bow[word] = idx
	}
	return BagOfWords{
		words: bow,
	}
}

func (bow BagOfWords) Contains(word string) bool {
	_, ok := bow.words[word]
	return ok
}

func TestInsert(t *testing.T) {
	trie := NewTrie()

	bow := NewBagOfWords("build", "builds")

	for word, testIdx := range bow.words {
		t.Logf("%d TestInsert: %s", testIdx, word)
		trie.Insert(word)
		node := trie.root
		for idx, char := range word {
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
	bow := NewBagOfWords("build", "builds")

	for word, testIdx := range bow.words {
		t.Logf("%d TestContains: %s", testIdx, word)
		trie.Insert(word)

		var path []rune
		for idx, char := range word {
			path = append(path, char)
			partialWord := string(path)
			if trie.Contains(partialWord) {
				if idx != len(word)-1 {
					if !bow.Contains(partialWord) {
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
