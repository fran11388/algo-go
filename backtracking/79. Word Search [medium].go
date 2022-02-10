package main

import "fmt"

func main() {
	board := make([][]byte, 3)
	board[0] = []byte("ABCE")
	board[1] = []byte("SFCS")
	board[2] = []byte("ADEE")
	word := "ABCB"
	fmt.Println(exist(board, word))
}

func exist(board [][]byte, word string) bool {
	// 嘗試每一點做backtrack 任一點可以  就回true
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			charIdx := 0
			can := backtrack(board, word, visited, r, c, charIdx)
			if can {
				return true
			}
		}
	}
	return false
}

func backtrack(board [][]byte, word string, visited [][]bool, r, c int, charIdx int) bool {
	if charIdx == len(word) {
		return true
	}

	isValid := isValidPos(board, word, visited, r, c, charIdx)
	if !isValid {
		return false
	}
	visited[r][c] = true

	//上
	if backtrack(board, word, visited, r-1, c, charIdx+1) {
		return true
	}

	//下
	if backtrack(board, word, visited, r+1, c, charIdx+1) {
		return true
	}

	//左
	if backtrack(board, word, visited, r, c-1, charIdx+1) {
		return true
	}

	//右
	if backtrack(board, word, visited, r, c+1, charIdx+1) {
		return true
	}
	visited[r][c] = false //退回
	return false
}

func isValidPos(board [][]byte, word string, visited [][]bool, r, c int, charIdx int) bool {
	if !(r >= 0 && r < len(board)) {
		return false
	}

	if !(c >= 0 && c < len(board[0])) {
		return false
	}

	if board[r][c] != word[charIdx] {
		return false
	}

	if visited[r][c] == true {
		return false
	}
	return true
}

//func exist(board [][]byte, word string) bool {
//	currStr:=""
//	tailPos:=[]int{}
//	return backtrack(board,word,&currStr,make(map[string]bool),&tailPos)
//}
//
//func backtrack(board [][]byte, word string, currStr *string, used map[string]bool,tailPos *[]int) bool {
//	//goal
//	if *currStr==word{
//		return true
//	}
//
//	//如果是第一個字的選擇  可以從board 任一點開始
//	if *currStr==""{
//		for row:=0;row<len(board);row++{
//			for col:=0;col<len(board[0]);col++{
//				ch:=board[row][col]
//
//				if ch==word[0]{//check constrain
//					*currStr=*currStr+string(ch)
//					key:=fmt.Sprintf("%d,%d",row,col)
//					used[key]=true
//					*tailPos=[]int{row,col}
//					isFind:=backtrack(board,word,currStr,used,tailPos)
//					if isFind{
//						return true
//					}
//					//undo choice
//					*currStr=""
//					delete(used,key)
//					*tailPos=[]int{}
//				}
//
//			}
//
//		}
//	}else{  //否則重先前的字開始延伸
//		choices :=getChoices(board,tailPos)
//		for _,choice:=range choices{
//
//			row:=choice[0]
//			col:=choice[1]
//			key:=fmt.Sprintf("%d,%d",row,col)
//			if _,ok:=used[key];ok{
//				continue
//			}
//			ch:=board[row][col]
//			currStrLen:=len(*currStr)
//			//constrain: //判斷該choice的字有沒有match 目標word
//			if ch!=word[currStrLen]{
//				continue
//			}
//
//			//make choice
//			*currStr=*currStr+string(ch)
//			used[key]=true
//			*tailPos=[]int{row,col}
//			isFind:=backtrack(board,word,currStr,used,tailPos)
//			if isFind{
//				return true
//			}
//			//undo choice
//			*currStr=(*currStr)[0:len(*currStr)-1]//移除尾部一個字
//			delete(used,key)
//			*tailPos=[]int{}
//		}
//
//	}
//
//
//	return false
//}
//
//func getChoices(board [][]byte,tailPos *[]int)[][]int{
//	result:=[][]int{}
//	row:=(*tailPos)[0]
//	col:=(*tailPos)[1]
//	maxRow:=len(board)
//	maxCol:=len(board[0])
//	//上
//	if row-1>=0{
//		result=append(result,[]int{row-1,col})
//	}
//
//	//下
//	if row+1<maxRow{
//		result=append(result,[]int{row+1,col})
//	}
//
//	//左
//	if col-1>=0{
//		result=append(result,[]int{row,col-1})
//	}
//
//	//右
//	if col+1<maxCol{
//		result=append(result,[]int{row,col+1})
//	}
//
//	return result
//}
