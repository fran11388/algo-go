package main

func searchMatrixByBinarySearch(matrix [][]int, target int) bool {
	size := len(matrix) * len(matrix[0])

	left := 0
	right := size - 1
	for left <= right && left >= 0 && right < size { //注意 左右指標重疊時也要做計算
		pivotIdx := (left + right) / 2
		r, c := getPos(matrix, pivotIdx)
		n := matrix[r][c]
		if n == target {
			return true
		} else if n > target {
			right = pivotIdx - 1
		} else {
			left = pivotIdx + 1
		}
	}
	return false
}

func getPos(matrix [][]int, pivotIdx int) (r, c int) {
	w := len(matrix[0])
	r = pivotIdx / w
	c = pivotIdx % w
	return r, c
}

func searchMatrixByTwoPointer(matrix [][]int, target int) bool {
	// init to bottom left
	r := len(matrix) - 1
	c := 0

	for r < len(matrix) && r >= 0 && c < len(matrix[r]) {
		n := matrix[r][c]

		if n == target {
			return true
		} else if n > target {
			r--
		} else {
			c++
		}
	}

	return false
}
