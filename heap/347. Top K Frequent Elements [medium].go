package main

import (
	"fmt"
	"math"
)

func main(){
	nums:=[]int{1,1,1,2,2,3}
	k:=2
	fmt.Println(topKFrequentByHeap(nums,k))
}

type Item struct{
	N int
	Freq int
}

func topKFrequentByHeap(nums []int, k int) []int {
	//init freq map
	freqMap:=map[int]int{}
	for _,n:=range nums{
		if _,ok:=freqMap[n];ok{
			freqMap[n]++
		}else{
			freqMap[n]=1
		}
	}

	minheap:=MinHeap{}
	for n,freq:=range freqMap{
		minheap.Insert(&Item{
			N:n,
			Freq: freq,
		})
		if len(minheap)>k{
			minheap.Remove()
		}
	}

	result:=[]int{}
	for len(minheap)>0{
		result=append(result,minheap.Remove().N)
	}
	return result
}


//time:O(n)
//space:O(n)
func topKFrequentByBucketSort(nums []int, k int) []int {
	//count frequent in hashmap
	count:=map[int]int{}
	for _,n:=range nums{
		if _,ok:=count[n];ok{
			count[n]++
		}else{
			count[n]=1
		}
	}


	//put the num in correspond frequent array
	freq:=make([][]int,len(nums)+1)
	for n,frequent:=range count{
		freq[frequent]=append(freq[frequent],n)
	}

	//iterate frequent array from tail and put the number into result

	result:=make([]int,0,k)
	for i:=len(freq)-1;i>=0;i--{
		for _,n:=range freq[i]{
			result=append(result,n)
			if len(result)==k{
				return result
			}
		}
	}
	return result
}

type MinHeap []*Item

func NewMinHeap(arr []*Item) *MinHeap {
	h := new(MinHeap)
	for _, n := range arr {
		h.Insert(n)
	}
	return h
}

func (h *MinHeap) Insert(n *Item) {
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
		if (*h)[idx].Freq < (*h)[parentIdx].Freq {
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

func (h *MinHeap) Remove() *Item {
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

func (h *MinHeap)removeLastElement()*Item{
	last:=(*h)[len(*h)-1]
	*h=(*h)[:len(*h)-1]
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
		if childIdx < 0 || (*h)[r].Freq < (*h)[childIdx].Freq {
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

	if h[left].Freq < h[right].Freq {
		return left
	} else {
		return right
	}
}

func (h MinHeap) Print() {
	fmt.Println(h)
}
