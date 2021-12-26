package main

import "fmt"

func main(){
	fmt.Println(canJump([]int{0,2,3}))
}

func canJump(nums []int) bool {
	return canJumpByTable(nums)
}

func canJumpHelp(nums []int, begin int, ht map[int]bool) bool {
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
			canjump := canJumpHelp(nums, begin+i, ht)
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
