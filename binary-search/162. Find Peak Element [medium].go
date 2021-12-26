package main

func findPeakElement(nums []int) int {
	return findPeakElementRecu(0,len(nums)-1,nums)
}

func findPeakElementRecu(left,right int ,nums []int) int {
	mid:=(left+right)/2
	if isPeak(nums,mid){
		return mid
	}

	if (mid+1)<len(nums) && nums[mid+1]> nums[mid]{
		return findPeakElementRecu(mid+1, right,nums)
	}else{
		return findPeakElementRecu(left, mid-1,nums)
	}

}



func isPeak(nums []int,p int)bool{
	leftIdx:=p-1
	rightIdx:=p+1

	if  leftIdx>=0 &&  nums[leftIdx]>=nums[p]{
		return false
	}

	if  rightIdx<len(nums) &&  nums[rightIdx]>=nums[p]{
		return false
	}

	return true
}
