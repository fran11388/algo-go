package main

import "fmt"

func main(){
	obj := Constructor(2)
	obj.Put(1,1)
	obj.Put(2,2)
	fmt.Println(obj.Get(1))
	obj.Put(3,3)
	fmt.Println(obj.Get(2))
}
type Node struct {
	Prev *Node
	Next *Node
	Val  int
	Key int
}

type LRUCache struct {
	dummyHead *Node
	dummyTail *Node
	size     int
	capacity int
	ht       map[int]*Node
}

func (l *LRUCache) prepend(n *Node) {
	n.Prev = l.dummyHead
	n.Next = l.dummyHead.Next

	//注意這邊很重要，要先暫存起來，才不會指錯(如果這邊不存tmp，也可以先將dummyHead的next的prev先指到新結點，然後才更新dummyHead的next到新節點)
	dummyheadNextTmp:=l.dummyHead.Next
	l.dummyHead.Next = n
	dummyheadNextTmp.Prev = n

	l.ht[n.Key]=n
	l.size++

	if l.size >l.capacity {
		//apply lru capacity
		l.remove(l.dummyTail.Prev)
	}
}

func (l *LRUCache) remove(n *Node) {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev

	n.Next=nil
	n.Prev=nil

	delete(l.ht,n.Key)
	l.size--
}

func (l *LRUCache) removeFirst() {
	l.remove(l.dummyHead.Next)
}

func Constructor(capacity int) LRUCache {
	h, t := &Node{}, &Node{}
	h.Next = t
	t.Prev = h
	return LRUCache{
		dummyHead: h,
		dummyTail: t,
		size:      0,
		capacity:  capacity,
		ht:        map[int]*Node{},
	}
}

func (this *LRUCache) Get(key int) int {
	if n, ok := this.ht[key]; ok {
		//apply lru policy
		this.remove(n)
		this.prepend(n)
		return n.Val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if n, ok := this.ht[key]; ok {
		//if exist, update value and move to first
		n.Val=value
		this.remove(n)
		this.prepend(n)
	} else {
		node:=&Node{Key:key,Val: value}
		this.prepend(node)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
