package main

import "fmt"

//每個位置有幾種到達方法 =上格+左格的方法數
//對於最邊界(最上邊 最左邊) 就只會有一種方式
func uniquePathsByBottomUp(m int, n int) int {
	dptable:=initTable(m,n)
	for c:=0;c<len(dptable[0]);c++{ //上邊初始化
		dptable[0][c]=1
	}

	for r:=0;r<len(dptable);r++{ //左邊界初始化
		dptable[r][0]=1
	}

	for r:=1;r<len(dptable);r++{
		for c:=1;c<len(dptable[r]);c++{
			//陣列從左而右，自上而下 依序累加 上格+左格
			dptable[r][c]= dptable[r-1][c]+dptable[r][c-1]
		}
	}
	r:=len(dptable)-1
	c:=len(dptable[0])-1
	return  dptable[r][c]

}

func initTable(m,n int)[][]int{
	r:=make([][]int,m)
	for i:=0;i<len(r);i++{
		r[i]=make([]int,n)
	}
	return r
}


//對於起點 計算往右+往下的走法
func uniquePathsByTopDown(m int, n int) int {
	ht := map[string]int{}
	return ways(m, n, 1, 1, ht)
}

func ways(m, n, r, c int, ht map[string]int) int {
	key := getKey(r, c)
	if v, ok := ht[key]; ok {
		return v
	}

	if r == m && c == n {
		return 1
	}

	if r > m || c > n {
		return 0
	}

	right := ways(m, n, r, c+1, ht)
	down := ways(m, n, r+1, c, ht)

	ht[key] = right + down
	return right + down
}

func getKey(r, c int) string {
	key := fmt.Sprintf("%d_%d", r, c)
	return key
}
