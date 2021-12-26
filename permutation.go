package main

import "fmt"

func main() {
	//fmt.Println("123")
	nums := []int{}
	for i := 1; i <= 3; i++ {
		nums = append(nums, i)
	}
	a:=permutationResu(nums)
	//a:=permutationIterative(nums)
	fmt.Println(a)
	//fmt.Println()
}

//自上而下 每一數都搭配子排列組合出新結果
func permutationResu(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{{nums[0]}}
	}

	var perms [][]int

	//loop take every nums
	//注意這個步驟
	//ht := make(map[int]bool)
	for idx, _ := range nums {
		var newnums []int
		takedNum := nums[idx]

		//if _, ok := ht[takedNum]; ok { //去重
		//	continue
		//}
		//ht[takedNum] = true

		newnums = append(newnums, nums[idx+1:]...)
		newnums = append(newnums, nums[0:idx]...)
		subPerms := permutationResu(newnums)
		for _, subperm := range subPerms {
			perm := append([]int{takedNum}, subperm...)
			perms = append(perms, perm)
		}
	}

	return perms
}

//自下而上
func permutationIterative(nums []int) [][]int {
	result := [][]int{{nums[0]}}

	for i := 1; i < len(nums); i++ {
		var newResult [][]int
		for _, perm := range result {
			numToAppend := nums[i]
			newperm := append([]int{numToAppend}, perm...)
			//note!!! avoid shallow copy
			newResult = append(newResult, append([]int{}, newperm...))
			//loop move position and deep copy append
			idx1 := 0
			for idx2 := 1; idx2 < len(newperm); idx2++ {
				swap(newperm, idx1, idx2)
				//note!!! avoid shallow copy
				newResult = append(newResult, append([]int{}, newperm...))
				idx1++
			}

		}
		//overwrite result
		result = newResult
	}
	return result
}

func swap(s []int, idx1, idx2 int) {
	t := s[idx1]
	s[idx1] = s[idx2]
	s[idx2] = t
}
