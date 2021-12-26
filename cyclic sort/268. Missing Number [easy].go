package main

//test case
//[9,6,4,2,3,5,7,0,1]
//[1,2]


// time:O(n)
// space:O(1)
//如果遍歷的數比array長度小，並且與遍歷的index不相等，那兩者互換，
//持續進行此步驟，直到不再滿足以上交換條件(該位置放的數=index or 該數=array長度)
func missingNumber(nums []int) int {

	for i,_:=range nums{
		// 注意 這邊要持續loop確保該位置放的數是正確的
		for true{
			j:=nums[i]
			if j!=len(nums) && i!=j{
				t:=nums[nums[i]]
				nums[nums[i]]=nums[i]
				nums[i]=t

			}else{
				break
			}
		}

	}

	for i,v:=range nums{

		if i!=v{
			return i
		}
	}
	return len(nums)

}