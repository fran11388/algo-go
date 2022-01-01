package main

import "fmt"

func main() {
	s := "dvdf"
	fmt.Println(lengthOfLongestSubstring(s))
}
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	l := 1
	left:= 0
	ch := string(s[0])
	charOccuIdxMap := map[string]int{ch: 0}
	for right := 1; right < len(s); right++ {
		ch := string(s[right])
		if idx, ok := charOccuIdxMap[ch]; ok {
			//若找到重複字  必須調整left ptr確保不重複，
			// 但有可能現在的left ptr已經超過重複字的第一個位置了，這時就不動；若還沒就移動至idx+1
			left = max(left, idx+1)

		}

		charOccuIdxMap[ch] = right

		if (right - left + 1) > l {
			l = right - left + 1
		}
	}
	return l
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
