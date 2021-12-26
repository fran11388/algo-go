package main

import "fmt"

func main() {
	nums := []int{5, 2, 6, 3, 8, 1, 9, 4, 7}
	fmt.Println(quickSort(nums))
}

func quickSort(nums []int) []int {
	return qSort(nums, 0, len(nums)-1)
}

func qSort(nums []int, start, end int) []int {
	if end<=start {
		return nums
	}

	//do partition
	pivot := partition(nums,start,end)
	//	sort left
	qSort(nums, start, pivot-1)
	//	sort right
	qSort(nums, pivot+1, end)

	return nums
}

func partition(nums []int,start, end int) int {

	//choos last ele as a pivot
	last := end
	pivot := nums[last]

	//iterate the prev elements ,compare with the pivot,
	//and maintain a pointer i, point to last element which less than the pivot
	i := start-1 //means at begin,no item less than pivot
	for j := start; j < last; j++ {
		if nums[j] > pivot { //if bigger than pivot, than no need for move i
			continue
		} else { //means find a number which less than pivot, so swap (i+1) and the j position
			tmp := nums[i+1] //the num bigger than pivot
			nums[i+1] = nums[j]
			nums[j] = tmp
			i++
		}

	}

	//and swap the pivot and (i+1) position
	tmp:=nums[i+1]
	nums[i+1]=nums[last]
	nums[last]=tmp
	return i+1
}
