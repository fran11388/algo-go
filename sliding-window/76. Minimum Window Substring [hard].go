package main

import "fmt"

func main() {
	s:="a"
	t:="a"
	fmt.Println(minWindow(s,t))
}
func minWindow(s string, t string) string {
	windowCharCount := map[string]int{}

	left, right := 0, 0
	minWindowsSize := len(s)
	minWindowLeft := -1
	minWindowRight := -1
	ch := string(s[0])
	windowCharCount[ch] = 1

	tCountMap:=map[string]int{}
	for i:=0;i<len(t);i++{
		c:=string(t[i])
		if _,ok:=tCountMap[c];ok{
			tCountMap[c]++
		}else{
			tCountMap[c]=1
		}
	}

	for left <= len(s) && right <= len(s) {
		if isDesirable(windowCharCount, t,tCountMap) {
			size := getWindowSize(left, right)
			if size <= minWindowsSize { //注意 這邊要是<=  因為minWindowsSize初始是設為s的長度
				minWindowsSize = size
				minWindowLeft = left
				minWindowRight = right
			}

			//remove char
			ch := string(s[left])
			windowCharCount[ch]--
			left++
			if left == len(s) {
				break
			}
		} else {
			//append char
			right++
			if right == len(s) {
				break
			}
			ch := string(s[right])
			if _, ok := windowCharCount[ch]; ok {
				windowCharCount[ch]++
			} else {
				windowCharCount[ch] = 1
			}

		}
	}
	if minWindowLeft == -1 {
		return ""
	}
	return string(s[minWindowLeft : minWindowRight+1])
}
func getWindowSize(left, right int) int {
	return right - left + 1
}

func isDesirable(windowCharCount map[string]int, t string,tCountMap map[string]int) bool {  //注意 這邊會有repeat string, 要比對t 每個char的數量

	for ch,count:=range tCountMap{
		if c,ok:=windowCharCount[ch];ok{
			if c<count{
				return false
			}
		}else{
			return false
		}
	}


	return true
}
