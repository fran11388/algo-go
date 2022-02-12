package main

import "fmt"

func main() {
	r := generateParenthesis(4)
	fmt.Println(r, len(r))
}

func generateParenthesis(n int) []string {
	result := []string{}
	currStr := ""
	backtrack(n, 0, 0, &currStr, &result)
	return result
}

func backtrack(n, open, close int, currStr *string, result *[]string) {
	if len(*currStr)==n*2{
		*result=append(*result,*currStr)
		return
	}

	if open<n{ //可放左括號
		*currStr=*currStr+"("
		backtrack(n,open+1,close,currStr,result)
		*currStr=(*currStr)[:len(*currStr)-1] //undo choice
	}

	if close<open{ //可放右括號
		*currStr=*currStr+")"
		backtrack(n,open,close+1,currStr,result)
		*currStr=(*currStr)[:len(*currStr)-1] //undo choice
	}
}


