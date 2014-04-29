package main

import (
	"fmt"
)

type Trie struct {
	isWord bool
	children []*Trie
}

func NewTrie() (*Trie) {
	return &Trie{false, make([]*Trie, 26)}
}

func main() {
	t := NewTrie
	fmt.Print(t)
}