package main

import "fmt"

func main() {
	case1 := []int{2,3,4,5,18,17,6}
	//fmt.Println(maxAreaBruteForce(case1))
	fmt.Println(maxAreaTwoPointer(case1))
}

func maxAreaTwoPointer(height []int) int {
	max := (len(height) - 1) * getmin(height[0], height[len(height)-1])
	l := 0
	r := len(height) - 1
	for l != r {
		area := (r - l) * getmin(height[l], height[r])
		if area > max {
			max = area
		}

		if height[l] < height[r] {  //注意 這邊要比柱子的高度
			l++
		} else {
			r--
		}
	}
	return max
}

func maxAreaBruteForce(height []int) int {
	max := 0

	for l := 0; l < len(height)-1; l++ {
		for r := l + 1; r < len(height); r++ {
			diff := r - l
			contain := diff * getmin(height[l], height[r]) //注意是要帶入高度
			if contain > max {
				max = contain
			}

		}
	}

	return max
}
func getmin(a, b int) int { //注意要回船2柱子最小的
	if a < b {
		return a
	}
	return b
}
