package main

import "fmt"

func main(){
	fmt.Println(canJump([]int{0,2,3}))
}

func canJump(nums []int) bool {
	return canJumpByTable(nums)
}

//從最後位置依序往左看，看能不能跳掉最後；若可以，更新last position，繼續往左看能不能跳到last position；
//若不行，就不更新last position，一樣繼續往左看能不能跳到last position
func canJumpByGreedy(nums []int) bool {
	if len(nums)<2{
		return true
	}
	lastPos:=len(nums)-1
	for i:=len(nums)-2;i>=0;i--{
		j:=nums[i]
		if i+j>=lastPos{
			lastPos=i
			if i==0{
				return true
			}
		}
	}
	return false
}

func canJumpRecu(nums []int, begin int, ht map[int]bool) bool {
	if begin == len(nums)-1 {
		return true
	}
	if a, ok := ht[begin]; ok {
		return a
	}

	//loop evert step, if any step can reach
	step := nums[begin]
	for i := 1; i <= step; i++ {
		if canjump,ok:=ht[begin+i];ok{
			if canjump{
				ht[begin] = true
				return true
			}
		}else{
			canjump := canJumpRecu(nums, begin+i, ht)
			ht[begin+i] = canjump
			if canjump {
				ht[begin] = true
				return true
			}
		}


	}

	ht[begin] = false
	return false
}

func canJumpByTable(nums []int) bool {
	table:=make([]bool,len(nums))
	table[0]=true
	for i:=0;i<len(nums)-1;i++{
		jump:=nums[i]
		if table[i]==false{ //若當前位置達不到  直接跳過不考慮
			continue
		}
		for j:=1;j<=jump;j++{
			if i+j<len(nums){
				table[i+j]=true
			}
		}
	}
	return table[len(nums)-1]
}
