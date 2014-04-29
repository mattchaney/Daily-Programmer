package main

import (
	"bufio"
	"fmt"
	"os"
)

type Trie struct {
	isWord bool
	children []*Trie
}

func NewTrie() (*Trie) {
	return &Trie{false, make([]*Trie, 26)}
}

func (t *Trie) insert(word string, pos int) {
	if pos == len(word) {
		t.isWord = true
	} else {
		chr := rune(word[pos]) - 'a'
		if t.children[chr] != nil {
			t.children[chr].insert(word, pos+1)
		} else {
			t.children[chr] = NewTrie()
			t.children[chr].insert(word, pos+1)
		}
	}
}

func (t *Trie) contains(word string, pos int) bool {
	if pos == len(word) {
		return t.isWord
	} else {
		chr := rune(word[pos]) - 'a'
		if t.children[chr] != nil {
			return t.children[chr].contains(word, pos+1)
		} else {
			return false
		}
	}
}

func load(t *Trie, scanner *bufio.Scanner) {
	for scanner.Scan() {
		str := scanner.Text()
		t.insert(str, 0)
	}
}

func addRune(recv, send string) (newRecv, newSend string) {
	newRecv = recv + string(send[0])
	newSend = send[1:]
	return
}

func combine(phrase *string, cons, vows string, results *[]string) {
	if len(cons) == 0 && len(vows) == 0 {
		*results = append(*results, *phrase)
	} else {
		if len(cons) > 0 {
			phrasePlusCon := new(string)
			var newCons string
			*phrasePlusCon, newCons = addRune(*phrase, cons)
			combine(phrasePlusCon, newCons, vows, results)
		}
		if len(vows) > 0 {
			phrasePlusVow := new(string)
			var newVows string
			*phrasePlusVow, newVows = addRune(*phrase, vows)
			combine(phrasePlusVow, cons, newVows, results)
		}
	}
}

func parsePhrase(curPhrase, curWord, remaining *string, results *[]string, root *Trie) {
	newCurWord := new(string)
	newRemaining := new(string)
	if len(*remaining) == 0 && root.contains(*curWord, 0) {
		*results = append(*results, *curPhrase)
	}
	*newCurWord, *newRemaining = addRune(*curWord, *remaining)
	if root.contains(*curWord, 0) {
		newCurPhrase := new(string)
		*newCurPhrase = *curPhrase + *newCurWord + " "
		newCurWord = new(string)
		parsePhrase(newCurPhrase, newCurWord, newRemaining, results, root)
		*newCurWord, *newRemaining = addRune(*curWord, *remaining)
		parsePhrase(newCurPhrase, newCurWord, newRemaining, results, root)
	}
}

func main() {
	file, _ := os.Open("enable1.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	
	root := NewTrie()
	load(root, scanner)

	word := os.Args[1]
	fmt.Println(word, "?", root.contains(word, 0))
	// var results []string
	// combine(new(string), os.Args[1], os.Args[2], &results)
	// // var validPhrases []string
	// for _, result := range results {
	// 	parsePhrase(new(string), new(string), &result, &validPhrases, root)
	// }
	// for valid := range validPhrases {
	// 	fmt.Println(valid)
	// }
}