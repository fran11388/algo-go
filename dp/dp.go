package main

import (
	"errors"
	"fmt"
)

func main() {

	//fmt.Println(gridTravelerRecu(5,5,make(map[string]int)))
	//fmt.Println(howSum(1,[]int{3,4}))
		fmt.Println(bestSum(10000000,[]int{2,3,5,7,1,2000,7999},make(map[int][]int)))
	//	str:="abcd"
	//fmt.Println(allConstruct("", []string{"purp","p","ur","le","prupl" }))

	//fmt.Println(fibTabulation(11))
	//fmt.Println(canSumTabulation(11,[]int{2,3,4}))
	//fmt.Println(howSumTabulation(60, []int{5,6,3}))
	//fmt.Println(bestSumTabulation(1000000,[]int{2,3,5,7,1,2000,7999,1000000}))
}





func gridTravelerRecu(x, y int, ht map[string]int) int {
	if x == 0 || y == 0 {
		return 0
	}
	if x == 1 && y == 1 {
		return 1
	}

	k := fmt.Sprintf("%d_%d", x, y)
	if v, ok := ht[k]; ok {
		return v
	}
	ht[k] = gridTravelerRecu(x-1, y, ht) + gridTravelerRecu(x, y-1, ht)

	return ht[k]
}
func canSum(n int, arr []int) bool {
	if n < 0 {
		return false
	}
	if n == 0 {
		return true
	}
	for _, v := range arr {
		remainder := n - v
		if canSum(remainder, arr) {
			return true
		}

	}
	return false
}

func howSum(n int, arr []int) ([]int, error) {
	if n == 0 {
		return []int{}, nil
	}
	if n < 0 {
		return nil, errors.New("not exist")
	}
	for _, v := range arr {
		remainder := n - v
		sum, err := howSum(remainder, arr)
		if err == nil {
			sum = append(sum, v)
			return sum, nil
		}

	}
	return nil, errors.New("not exist")
}

func bestSum(n int, arr []int,ht map[int][]int) ([]int, error) {
	if result,ok:=ht[n];ok{
		return result,nil
	}
	if n == 0 {
		return []int{}, nil
	}
	if n < 0 {
		return nil, errors.New("not exist")
	}

	var best []int
	isExist := false
	for _, v := range arr {
		remainder := n - v
		sum, err := bestSum(remainder, arr,ht)
		if err == nil {
			sum = append(sum, v)
			if isExist {
				if len(sum) < len(best) {
					best = sum
				}
			} else {
				isExist = true
				best = sum
			}
		}

	}
	if isExist {
		ht[n]=best
		return best, nil
	}
	return nil, errors.New("not exist")
}

func fibTabulation(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	result := make([]int, n)
	result[0] = 0
	result[1] = 1
	for i := 0; i < n; i++ {
		nextOne := i + 1
		nextTwo := i + 2
		if nextOne < n {
			result[nextOne] = result[nextOne] + result[i]
		}
		if nextTwo < n {
			result[nextTwo] = result[nextTwo] + result[i]
		}
	}
	return result[n-1]
}

func canSumTabulation(n int, arr []int) bool {
	t := make([]bool, n+1)
	t[0] = true
	for idx, canSum := range t {
		if canSum {
			for _, num := range arr {
				target := idx + num
				if target <= n {
					t[target] = true
				}
			}
		}

	}
	return t[n]
}

func howSumTabulation(targetSum int, numbers []int) []int {
	table := make([][]int, targetSum+1)
	table[0] = []int{}
	for currNum, t := range table {
		if t != nil {
			for _, number := range numbers {
				target := currNum + number
				if target<=targetSum{
					newslice:=make([]int,len(t)+1)
					copy(newslice,t)
					newslice[len(newslice)-1]=number
					table[target]=newslice
				}
			}
		}
	}
	return table[targetSum]
}

func bestSumTabulation(targetSum int, numbers []int)[]int{

	//numbers
	//[1,2,3]

	//table
	//0   1     2    3      4       5
	//___________________________________
	//[]  [1]  [2]  [3]   [3,1]   [2,3]
	table :=make([][]int,targetSum+1)
	table[0]=[]int{}

	for tableIdx, tableWays :=range table{
		if tableWays !=nil{
			for _,number:=range numbers{
				newSum:= tableIdx +number
				if newSum<=targetSum{
					newWays:=append(append([]int{}, tableWays...),number)
					if len(newWays)<len(table[newSum])&&table[newSum]!=nil{
						table[newSum]=newWays
					}else if table[newSum]==nil{
						table[newSum]=newWays
					}
				}
			}
		}

	}

	return table[targetSum]
}
