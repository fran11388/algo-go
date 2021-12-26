package main

import "fmt"

func main() {
	interval := [][]int{{1,2}, {3,5},{6,7},{8,10},{12,16}}
	newInterval := []int{4,8}
	fmt.Println(insert(interval, newInterval))
}

func insert(intervals [][]int, newInterval []int) [][]int {

	result := [][]int{}
	isnewIntervalBeenAdded := false
	for _, interval := range intervals {
		if canMerge(newInterval, interval) {
			mergeToNewInterval(newInterval, interval)
			continue
		} else if isNewIntervalNeedAddtoResult(newInterval, interval) && !isnewIntervalBeenAdded {
			result = append(result, []int{newInterval[0], newInterval[1]})
			isnewIntervalBeenAdded = true
		}

		if isnewIntervalBeenAdded {
			result = append(result, []int{interval[0], interval[1]})
		}
	}

	return result
}

func isNewIntervalNeedAddtoResult(newInterval []int, interval []int) bool {
	return newInterval[1] < interval[0]
}

func canMerge(interval []int, currInterval []int) bool {
	left := interval[0]
	right := interval[1]
	currLeft := currInterval[0]
	currRight := currInterval[1]

	if currRight <= right && currRight >= left {
		return true
	}

	if currLeft >= left && currLeft <= right {
		return true
	}
	if currLeft <= left && currRight >= right {
		return true
	}

	return false
}

func mergeToNewInterval(newInterval []int, interval []int) {
	newInterval[0] = getMin(interval[0], newInterval[0])
	newInterval[1] = getMax(interval[1], newInterval[1])
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
