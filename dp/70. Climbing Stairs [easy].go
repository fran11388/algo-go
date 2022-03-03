package main

func climbStairs(n int) int {
	table := make([]int, n+1)
	table[0] = 1
	//考慮到第n-1
	for i := 0; i < n; i++ {
		if i+1 <= n {
			table[i+1] += table[i]
		}
		if i+2 <= n {
			table[i+2] += table[i]
		}
	}
	return table[n]
}
