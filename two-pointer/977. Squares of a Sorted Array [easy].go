package main
func sortedSquares(nums []int) []int {
	result:=make([]int,len(nums))
	appendIdx:=len(nums)-1

	left:=0
	right:=len(nums)-1
	for left<=right{
		n1:=nums[left]*nums[left]
		n2:=nums[right]*nums[right]

		if n1>n2{
			result[appendIdx]=n1
			left++
		}else{
			result[appendIdx]=n2
			right--
		}
		appendIdx--
	}
	return result
}