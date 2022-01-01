package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aab"
	fmt.Println(reorganizeString(s))
}

func reorganizeString(s string) string {
	res := ""
	ht := map[string]int{}
	maxCountChar := ""
	maxCount := 0
	for i := 0; i < len(s); i++ {
		ch := string(s[i])
		if _, ok := ht[ch]; ok {
			ht[ch]++

		} else {
			ht[ch] = 1
		}

		if ht[ch] > maxCount {
			maxCountChar = ch
			maxCount = ht[ch]
		}
	}

	charList := make([]string, len(s))
	putIdx := 0
	for i := 0; i < ht[maxCountChar]; i++ {
		if putIdx>=len(charList){  //max freq 的自若太多 案偶數置放會超出array size
			return ""
		}
		charList[putIdx] = maxCountChar
		putIdx += 2
	}

	nextPutPos := 0
	for ch, count := range ht {
		if ch == maxCountChar {
			continue
		}

		for i := 0; i < count; i++ {
			for true { //若放的位置不為空  向後移動
				if charList[nextPutPos] != "" {
					nextPutPos++
				} else {
					break
				}
			}

			charList[nextPutPos] = ch
			nextPutPos++
		}
	}

	res = strings.Join(charList, "")
	return res
}
