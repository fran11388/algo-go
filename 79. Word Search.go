package main

import "fmt"

func main() {
	board := make([][]byte, 3)
	board[0] = []byte("ABCE")
	board[1] = []byte("SFCS")
	board[2] = []byte("ADEE")
	word := "ABCCED"



	fmt.Println(exist(board, word))
}

func exist(board [][]byte, word string) bool {
	currStr:=""
	tailPos:=[]int{}
	return backtrack(board,word,&currStr,make(map[string]bool),&tailPos)
}

func backtrack(board [][]byte, word string, currStr *string, used map[string]bool,tailPos *[]int) bool {
	//goal
	if *currStr==word{
		return true
	}

	//如果是第一個字的選擇  可以從board 任一點開始
	if *currStr==""{
		for row:=0;row<len(board);row++{
			for col:=0;col<len(board[0]);col++{
				ch:=board[row][col]

				if ch==word[0]{//check constrain
					*currStr=*currStr+string(ch)
					key:=fmt.Sprintf("%d,%d",row,col)
					used[key]=true
					*tailPos=[]int{row,col}
					isFind:=backtrack(board,word,currStr,used,tailPos)
					if isFind{
						return true
					}
					//undo choice
					*currStr=""
					delete(used,key)
					*tailPos=[]int{}
				}

			}

		}
	}else{  //否則重先前的字開始延伸
		choices :=getChoices(board,tailPos)
		for _,choice:=range choices{

			row:=choice[0]
			col:=choice[1]
			key:=fmt.Sprintf("%d,%d",row,col)
			if _,ok:=used[key];ok{
				continue
			}
			ch:=board[row][col]
			currStrLen:=len(*currStr)
			//constrain: //判斷該choice的字有沒有match 目標word
			if ch!=word[currStrLen]{
				continue
			}

			//make choice
			*currStr=*currStr+string(ch)
			used[key]=true
			*tailPos=[]int{row,col}
			isFind:=backtrack(board,word,currStr,used,tailPos)
			if isFind{
				return true
			}
			//undo choice
			*currStr=(*currStr)[0:len(*currStr)-1]//移除尾部一個字
			delete(used,key)
			*tailPos=[]int{}
		}

	}


	return false
}

func getChoices(board [][]byte,tailPos *[]int)[][]int{
	result:=[][]int{}
	row:=(*tailPos)[0]
	col:=(*tailPos)[1]
	maxRow:=len(board)
	maxCol:=len(board[0])
	//上
	if row-1>=0{
		result=append(result,[]int{row-1,col})
	}

	//下
	if row+1<maxRow{
		result=append(result,[]int{row+1,col})
	}

	//左
	if col-1>=0{
		result=append(result,[]int{row,col-1})
	}

	//右
	if col+1<maxCol{
		result=append(result,[]int{row,col+1})
	}

	return result
}

