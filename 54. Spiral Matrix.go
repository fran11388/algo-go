package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralOrder(matrix))
}

func spiralOrder(matrix [][]int) []int {
	result := []int{}
	totalNum := len(matrix) * len(matrix[0])

	rightmostX := len(matrix[0]) - 1
	downMostY := len(matrix) - 1
	leftMostX := 0
	upMostY := 0
	for len(result) < totalNum {
		//向右
		y := upMostY
		x := leftMostX
		for x <= rightmostX {
			result = append(result, matrix[y][x])
			if len(result) == totalNum{
				return result
			}
			x++
		}
		upMostY++

		//向下
		x = rightmostX
		y= upMostY
		for y <= downMostY {
			result = append(result, matrix[y][x])
			if len(result) == totalNum{
				return result
			}
			y++
		}
		rightmostX--

		//向左
		y = downMostY
		x--
		for x >= leftMostX {
			result = append(result, matrix[y][x])
			if len(result) == totalNum{
				return result
			}
			x--
		}
		downMostY--

		//向上
		x = leftMostX
		y--
		for y >= upMostY {
			result = append(result, matrix[y][x])
			if len(result) == totalNum{
				return result
			}
			y--
		}
		leftMostX++
	}

	return result
}
