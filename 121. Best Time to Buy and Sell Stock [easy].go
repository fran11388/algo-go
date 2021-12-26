package main

import "fmt"

func main() {
	p := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(p))
}

func maxProfit(prices []int) int {
	min := 100000
	maxprofit := 0
	for _, p := range prices {
		if p < min {
			min = p
		}

		if p-min > maxprofit {
			maxprofit = p - min
		}
	}
	return maxprofit
}
