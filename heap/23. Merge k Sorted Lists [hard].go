package main

import (
	"fmt"
	"math"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var minheap MinHeap
	for _, node := range lists {
		if node != nil {
			minheap.Insert(node)
		}
	}

	var r *ListNode
	if len(minheap) == 0 {
		return r
	}

	popedNode := minheap.Remove()
	r = popedNode
	root:=r
	if popedNode.Next != nil {
		minheap.Insert(popedNode.Next)
	}

	for len(minheap) > 0 {
		popedNode := minheap.Remove()
		r.Next = popedNode
		r=popedNode
		if popedNode.Next != nil {
			minheap.Insert(popedNode.Next)
		}
	}

	r.Next = nil
	return root
}

type MinHeap []*ListNode

func NewMinHeap(arr []*ListNode) *MinHeap {
	h := new(MinHeap)
	for _, n := range arr {
		h.Insert(n)
	}
	return h
}

func (h *MinHeap) Insert(n *ListNode) {
	*h = append(*h, n)
	h.bubbleUp(len(*h) - 1)
}
func (h *MinHeap) getInsertPosition() int {
	return len(*h)
}

func (h *MinHeap) getParent(idx int) int {
	return int(math.Floor(float64(idx-1) / float64(2)))
}

func (h *MinHeap) bubbleUp(idx int) {
	for idx != 0 {
		parentIdx := h.getParent(idx)
		if (*h)[idx].Val < (*h)[parentIdx].Val {
			h.swap(idx, parentIdx)
			idx = parentIdx
		} else {
			break
		}
	}
}

func (h *MinHeap) swap(idx1, idx2 int) {
	t := (*h)[idx1]
	(*h)[idx1] = (*h)[idx2]
	(*h)[idx2] = t

}

func (h *MinHeap) Remove() *ListNode {
	if len(*h) == 0 {
		return nil
	}
	if len(*h) == 1 {
		return h.removeLastElement()
	}

	peak := (*h)[0]
	h.swapLastWithRoot()
	h.removeLastElement()
	h.topDown()

	return peak
}

func (h *MinHeap) removeLastElement() *ListNode {
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

func (h *MinHeap) swapLastWithRoot() {
	last := len(*h) - 1
	h.swap(0, last)
}

func (h *MinHeap) topDown() {
	r := 0
	for true {
		childIdx := h.getMinimumChildIdx(r)
		if childIdx < 0 || (*h)[r].Val < (*h)[childIdx].Val {
			break
		}
		h.swap(r, childIdx)
		r = childIdx
	}

}
func (h MinHeap) getMinimumChildIdx(n int) int {
	left := (n * 2) + 1
	right := (n * 2) + 2
	idxLeng := len(h) - 1

	if left > idxLeng {
		return -1
	} else if right > idxLeng {
		return left
	}

	if h[left].Val < h[right].Val {
		return left
	} else {
		return right
	}
}

func (h MinHeap) Print() {
	fmt.Println(h)
}
