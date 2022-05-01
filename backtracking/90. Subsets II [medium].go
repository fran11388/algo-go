package main

func subsetsWithDup(nums []int) [][]int {
	result := [][]int{}
	elements := []int{}
	idx := 0
	notConsiderNums := map[int]bool{}
	backtrack(nums, idx, notConsiderNums, &elements, &result)
	return result
}

//思路: 對於每一elemt有考慮與不考慮，但因為input有重複元素，
//所以在做了不考慮某數後，往後做的其他選擇也必須不考慮該數；否則會有重複子集
func backtrack(nums []int, idx int, notConsiderNums map[int]bool, elements *[]int, result *[][]int) {
	if idx == len(nums) {
		r := make([]int, len(*elements))
		copy(r, *elements)

		*result = append(*result, r)
		return
	}

	n := nums[idx]
	if _, ok := notConsiderNums[n]; ok {
		backtrack(nums, idx+1, notConsiderNums, elements, result)
		return
	}

	//考慮n
	*elements = append(*elements, n)
	backtrack(nums, idx+1, notConsiderNums, elements, result)
	*elements = (*elements)[:len(*elements)-1]

	//不考慮n
	notConsiderNums[n] = true
	backtrack(nums, idx+1, notConsiderNums, elements, result)
	delete(notConsiderNums, n)
}
