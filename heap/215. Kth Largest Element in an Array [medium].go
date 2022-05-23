package main

import (
	"fmt"
	"github.com/emirpasic/gods/trees/binaryheap"
)

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargestByQuickSelect(nums, 2))
}

//kth largest會在n-k的位置(小到大)
//不斷找pivot直到pivot位在n-k
func findKthLargestByQuickSelect(nums []int, k int) int {
	targetIdx := len(nums) - k
	start, end := 0, len(nums)-1
	for true {
		pivotIdx := findPivot(nums, start, end)
		if pivotIdx == targetIdx {
			return nums[pivotIdx]
		} else if targetIdx > pivotIdx {
			start = pivotIdx + 1
		} else {
			end = pivotIdx - 1
		}
	}
	return -1
}

func findPivot(nums []int, start, end int) int {
	randomPivot := end
	i, j := start-1, start
	for ; j < randomPivot; j++ {
		n := nums[j]
		if n < nums[randomPivot] {
			nums[i+1], nums[j] = nums[j], nums[i+1]
			i++
		}
	}

	nums[i+1], nums[randomPivot] = nums[randomPivot], nums[i+1]
	return i + 1
}

func findKthLargestByHeap(nums []int, k int) int {
	heap := binaryheap.NewWith(func(a, b interface{}) int {
		va := a.(int)
		vb := b.(int)
		return va - vb
	})

	for _, n := range nums {
		heap.Push(n)

		if heap.Size() > k {
			heap.Pop()
		}
	}
	n, _ := heap.Pop()
	return n.(int)
}
