package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Entries struct {
	Entries []Entry `json:"entries"`
}

type Entry struct {
	Word       string `json:"word"`
	Pos        string `json:"type"`
	Definition string `json:"description"`
}

type Definition struct {
	partOfSpeech string
	definition   string
}

type Dictionary struct {
	words map[string][]Definition
}

type WordsTable struct {
	ID   uint
	Word string
}

func NewDictionary() Dictionary {
	return Dictionary{
		words: make(map[string][]Definition),
	}
}

func (dict *Dictionary) Insert(entry Entry) {
	definition := Definition{
		partOfSpeech: entry.Pos,
		definition:   entry.Definition,
	}
	dict.words[entry.Word] = append(dict.words[entry.Word], definition)
}

func (dict *Dictionary) Count() int {
	return len(dict.words)
}

func (dict *Dictionary) List() []string {
	var wordList []string

	for word := range dict.words {
		wordList = append(wordList, word)
	}
	return wordList
}

func (dict *Dictionary) Read(word string) {
	fmt.Println(word)
	for _, def := range dict.words[word] {
		fmt.Println(def.partOfSpeech, def.definition)
	}
}

func readJSONDictionary(dictionaryFile string) Dictionary {
	jsonFile, err := os.Open(dictionaryFile)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var entries Entries
	json.Unmarshal(byteValue, &entries)

	dict := NewDictionary()
	for _, entry := range entries.Entries {
		dict.Insert(entry)
	}
	return dict
}

func test_dict() {
	dict := readJSONDictionary("dictionary.json")
	// fmt.Println(dict.Count())
	// fmt.Println(dict.List())
	dict.Read("Canter")

}
