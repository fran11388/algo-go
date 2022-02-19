package main

func combinationSum3(k int, n int) [][]int {
	beginNum := 1
	currSum := 0
	result := [][]int{}
	nums := []int{}
	backtrack(k, n, beginNum, currSum, &nums, &result)
	return result
}

func backtrack(k int, n int, beginNum int, currSum int, nums *[]int, result *[][]int) {
	if len(*nums) == k && currSum == n {
		newnums := make([]int, k)
		copy(newnums, *nums)
		*result = append(*result, newnums)
		return
	}

	if currSum > n && len(*nums) > k {
		return
	}

	for i := beginNum; i <= 9; i++ {
		*nums = append(*nums, i)                      //make choice
		backtrack(k, n, i+1, currSum+i, nums, result) //因為每個number最多用一次, 所以下次考慮時，必須是當前選的數+1(i+1)
		*nums = (*nums)[:len(*nums)-1]                //undo choice
	}
}
