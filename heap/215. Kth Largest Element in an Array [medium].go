package main

import (
	"fmt"
	"math"
)

func main(){
	nums:=[]int{3,2,1,5,6,4}
	fmt.Println(findKthLargest(nums,2))
}

func findKthLargest(nums []int, k int) int {
	heap := &MinHeap{}
	for _, n := range nums {
		heap.Insert(n)  //注意  這邊要先insert，有超過才remove，這樣remove掉的才會是小的
		if len(*heap) > k {
			heap.Remove()
		}
	}

	return (*heap)[0]
}

type MinHeap []int

func NewMinHeap(arr []int) *MinHeap {
	h := new(MinHeap)
	for _, n := range arr {
		h.Insert(n)
	}
	return h
}

func (h *MinHeap) Insert(n int) {
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
		if (*h)[idx] < (*h)[parentIdx] {
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

func (h *MinHeap) Remove() int {
	if len(*h) == 0 {
		return -1
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

func (h *MinHeap) removeLastElement() int {
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
		if childIdx < 0 || (*h)[r] < (*h)[childIdx] {
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

	if h[left] < h[right] {
		return left
	} else {
		return right
	}
}
