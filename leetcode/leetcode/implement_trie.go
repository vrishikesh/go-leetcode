package leetcode

import "fmt"

type Trie interface {
	insert(word string)
	search(word string) bool
	startsWith(prefix string) bool
}

type trieNode struct {
	value   string
	chidren map[rune]*trieNode
	isEnd   bool
}

type trie struct {
	root *trieNode
}

func NewTrie() *trie {
	return &trie{
		root: &trieNode{},
	}
}
func (t *trie) insert(word string) {
	node := t.root
	for _, r := range word {
		idx := r - 'a'
		if node.chidren[idx] == nil {
			node.chidren[idx] = &trieNode{value: string(r)}
		}
		node = node.chidren[idx]
	}
	node.isEnd = true
}
func (t *trie) search(word string) bool {
	node := t.root
	for _, r := range word {
		idx := r - 'a'
		if node.chidren[idx] == nil {
			return false
		}
		node = node.chidren[idx]
	}
	return node.isEnd
}
func (t *trie) startsWith(prefix string) bool {
	node := t.root
	for _, r := range prefix {
		idx := r - 'a'
		if node.chidren[idx] == nil {
			return false
		}
		node = node.chidren[idx]
	}
	return true
}

func ImplementTrie() {
	t := NewTrie()
	t.insert("apple")
	fmt.Println(t.search("dog")) // false
	t.insert("dog")
	fmt.Println(t.search("dog"))     // true
	fmt.Println(t.startsWith("app")) // true
	fmt.Println(t.search("app"))     // false
	t.insert("app")
	fmt.Println(t.search("app")) // true

	var printTrie func(node *trieNode)
	printTrie = func(node *trieNode) {
		for _, c := range node.chidren {
			if c != nil {
				fmt.Println(c.value, c.isEnd)
				printTrie(c)
			}
		}
	}
	printTrie(t.root)
}
