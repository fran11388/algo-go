package main

func letterCombinations(digits string) []string {
	var result []string
	digitMpa:=map[byte][]string{
		'2':[]string{"a","b","c"},
		'3':[]string{"d","e","f"},
		'4':[]string{"g","h","i"},
		'5':[]string{"j","k","l"},
		'6':[]string{"m","n","o"},
		'7':[]string{"p","q","r","s"},
		'8':[]string{"t","u","v"},
		'9':[]string{"w","x","y","z"},

	}
	currStr:=""
	backtracking(digits,0,&currStr,&result,digitMpa)
	return  result
}


func backtracking(digits string,currIdx int, currStr *string,result *[]string,digitMpa map[byte][]string){
	if digits==""{
		return
	}

	if currIdx>=len(digits) &&(*currStr)!=""{ //空字串不放入
		*result=append(*result,*currStr)
		return
	}

	currChar:=digits[currIdx]
	charlist:=digitMpa[currChar]
	for _,c:=range charlist{
		*currStr=(*currStr)+c //make choice
		backtracking(digits,currIdx+1,currStr,result, digitMpa )

		//undo choice: remove last char
		currStrLen:=len(*currStr)
		*currStr=(*currStr)[:currStrLen-1]

	}
}
