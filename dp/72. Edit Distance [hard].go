package main

func minDistance(word1 string, word2 string) int {
	table:=make([][]int,len(word2)+1)
	for i:=0;i<=len(word2);i++{
		table[i]=make([]int,len(word1)+1)
	}


	// init table
	for i:=0;i<=len(word1);i++{
		table[0][i]=i //delete operation
	}
	for i:=0;i<=len(word2);i++{
		table[i][0]=i//insert operation
	}


	// iterate throughout table by bottom up
	for r:=1;r<=len(word2);r++{
		for c:=1;c<=len(word1);c++{
			charA:=getCharByIdx(word1,c-1)
			charB:=getCharByIdx(word2,r-1)
			if charA==charB{
				table[r][c]=table[r-1][c-1] //代表不用編輯
			}else{
				cell1:=table[r-1][c-1]//replace
				cell2:=table[r-1][c]//delete
				cell3:=table[r][c-1]//insert
				table[r][c]=getMin(cell1,cell2,cell3)+1
			}

		}
	}

	r:=len(word2)
	c:=len(word1)
	return table[r][c]
}

func getCharByIdx(str string,i int)byte{
	return str[i]
}

func getMin(p ...int)int{
	m:=p[0]
	for _,v:=range p{
		if v<m{
			m=v
		}
	}
	return m
}