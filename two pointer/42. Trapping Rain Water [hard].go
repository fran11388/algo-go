package main

import "fmt"

func main() {
	t := []int{2, 1, 0, 1, 2}
	fmt.Println(trapByMonotonicStack(t))
}

func trapByMonotonicStack(height []int) int {
	var stack []int
	result := 0
	for currIdx, _ := range height {
		if len(stack) == 0 {
			stack = append(stack, currIdx)
			continue
		}

		currH := height[currIdx]
		for true {
			if currH > stackTopHeight(stack, height) {
				popedIdx := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				popedH := height[popedIdx]
				if len(stack)==0{
					break
				}
				stH := stackTopHeight(stack, height)
				stIdx := stackTopIdx(stack)
				h := getMin(currH, stH) - popedH
				d := currIdx - stIdx - 1
				if h > 0 && stIdx>=0{
					w := d * h
					if w > 0 {
						result += w
					}

				}
				continue
			}
			break
		}

		stack = append(stack, currIdx)
	}
	return result

}
func stackTopHeight(stack []int, height []int) int {
	if len(stack) == 0 {
		return 0
	}
	idx := stack[len(stack)-1]
	return height[idx]
}

func stackTopIdx(stack []int)int{
	if len(stack) == 0 {
		return -1
	}
	idx := stack[len(stack)-1]
	return idx

}

func trapByTwoPointer(height []int) int {
	result := 0
	maxL := height[0]
	maxR := height[len(height)-1]
	l := 0
	r := len(height) - 1

	for l != r {
		if maxL <= maxR {
			l++
			currHeight := height[l]
			if w := getMin(maxL, maxR) - currHeight; w > 0 {
				result += w
			}
			if currHeight > maxL {
				maxL = currHeight
			}
		} else {
			r--
			currHeight := height[r]
			if w := getMin(maxL, maxR) - currHeight; w > 0 {
				result += w
			}
			if currHeight > maxR {
				maxR = currHeight
			}
		}
	}

	return result
}

func trapDP(height []int) int {
	currMaxHeightFromLeft := make([]int, len(height))
	maxleft := 0
	for idx, h := range height {
		currMaxHeightFromLeft[idx] = maxleft
		if h > maxleft {
			maxleft = h
		}
	}

	currMaxHeightFromRight := make([]int, len(height))
	maxRight := 0
	for i := len(height) - 1; i >= 0; i-- {
		currMaxHeightFromRight[i] = maxRight
		if height[i] > maxRight {
			maxRight = height[i]
		}
	}

	result := 0
	for idx, h := range height {
		trapped := getMin(currMaxHeightFromLeft[idx], currMaxHeightFromRight[idx]) - h
		if trapped > 0 {
			result += trapped
		}
	}

	//fmt.Println(currMaxHeightFromLeft)
	return result
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
