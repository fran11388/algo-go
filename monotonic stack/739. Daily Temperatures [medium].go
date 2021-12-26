package monotonic_stack

//use monotonic stack
func dailyTemperatures(temperatures []int) []int {
	out := make([]int, len(temperatures))
	stack := [][]int{}
	for idx, temp := range temperatures {
		if len(stack) == 0 {
			stack = append(stack, []int{idx, temp})
		} else {
			//把當前的溫度放入stack 並需要維持montonic
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				topIdx := top[0]
				topTemp := top[1]

				//如果當前溫度大於top stack，表示找到warmer day 計算並pop up
				if temp > topTemp {
					// calculate warmer day
					out[topIdx] = idx - topIdx
					// pop
					stack = stack[:len(stack)-1]
				} else {
					//如果沒有表示可以放入stack 直接跳出迴圈
					break
				}
			}
			//把當前的溫度放入stack
			stack = append(stack, []int{idx, temp})
		}
	}
	return out
}
