package main

import "fmt"

func main() {
	r:=generateParenthesis(4)
	fmt.Println(r,len(r))
}

func generateParenthesis(n int) []string {
	return generateParenthesisHelper("",n,n)
}
func generateParenthesisHelper(currStr string, open int, close int) []string {
	if open == 0 && close == 0 {
		return []string{currStr}
	}
	var result []string
	var leftResult []string
	var rightResult []string
	//放左括號
	if open > 0 {
		appendLeftBracket := currStr + "("
		leftResult = generateParenthesisHelper(appendLeftBracket, open-1, close)
	}

	//放右括號
	if close > 0 && close > open {
		appendRightBracket := currStr + ")"
		rightResult = generateParenthesisHelper(appendRightBracket, open, close-1)
	}

	//回傳左右的加總
	result = append(result, leftResult...)
	result = append(result, rightResult...)
	return result
}
