package main

func generate(numRows int) [][]int {
	result := make([][]int, 0, numRows)
	result = append(result, []int{1})

	for i := 2; i <= numRows; i++ {
		row := make([]int, 0, i)
		previousRow := result[i-2]
		for j := 0; j < i; j++ {
			prevLeftIdx := j - 1
			prevRightIdx := j
			node := 0
			if prevLeftIdx >= 0 {
				node += previousRow[prevLeftIdx]
			}

			if prevRightIdx < len(previousRow) {
				node += previousRow[prevRightIdx]
			}
			row = append(row, node)

		}

		result = append(result, row)
	}
	return result

}
