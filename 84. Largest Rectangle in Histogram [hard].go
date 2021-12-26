package main

import "fmt"

func main(){
	heights := []int{2,1,5,6,2,3}
	fmt.Println(largestRectangleAreaBruteForce(heights))
}
func largestRectangleAreaBruteForce(heights []int) int {
	max:=0
	for i:=0;i<len(heights);i++{
		for j:=i;j<len(heights);j++{
			area:=(j-i+1)*getMin(heights[i:j+1])
			if area>max{
				max=area
			}
		}
	}
	return max
}

func getMin(arr []int)int{
	if len(arr)<2{
		return arr[0]
	}
	min:=arr[0]
	for _,v:=range arr{
		if v<min{
			min=v
		}
	}
	return min
}