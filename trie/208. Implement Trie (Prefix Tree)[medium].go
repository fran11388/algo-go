package main

import "fmt"

func main() {
	str := "apple"
	t := Constructor()
	t.Insert(str)
	fmt.Println(t.Search("app"))
	fmt.Println(t.Search("apple"))
	fmt.Println(t.StartsWith("app"))
	fmt.Println(t.StartsWith("apple"))

}

type TrieNode struct {
	Val       string
	Child     map[string]*TrieNode
	EndOfWord bool
}

type Trie struct {
	Root *TrieNode
}

func Constructor() Trie {
	return Trie{Root: &TrieNode{Child: map[string]*TrieNode{}}}
}

func (this *Trie) appendNodeIfNotExist(node *TrieNode, char string) *TrieNode {
	if n, ok := node.Child[char]; ok {
		return n
	} else {
		newnode := &TrieNode{
			Child: map[string]*TrieNode{},
		}
		node.Child[char] = newnode
		return newnode
	}
}

func (this *Trie) Insert(word string) {
	chars := getChars(word)
	nodePtr := this.Root
	for _, c := range chars {
		nodePtr = this.appendNodeIfNotExist(nodePtr, c)
	}
	nodePtr.EndOfWord = true
}
func getChars(str string) []string {
	var arr []string
	for i := 0; i < len(str); i++ {
		arr = append(arr, string(str[i]))

	}
	return arr
}

func (this *Trie) findChild(node *TrieNode, char string) *TrieNode {
	if n, ok := node.Child[char]; ok {
		return n
	}
	return nil
}

func (this *Trie) Search(word string) bool {
	node:=this.startsWithAndGetNode(word)

	if node==nil{
		return false
	}

	if node.EndOfWord == false {
		return false
	}

	return true
}

func (this *Trie) StartsWith(prefix string) bool {
	if this.startsWithAndGetNode(prefix) != nil {
		return true
	} else {
		return false
	}
}

func (this *Trie) startsWithAndGetNode(prefix string) *TrieNode {
	nodePtr := this.Root
	chars := getChars(prefix)
	for _, c := range chars {
		nodePtr = this.findChild(nodePtr, c)
		if nodePtr == nil {
			return nil
		}
	}

	return nodePtr
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
