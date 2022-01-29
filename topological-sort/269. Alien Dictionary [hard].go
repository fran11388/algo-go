package main

import "fmt"

//coding 55分完成  一次就過lol
func alienOrder(words []string) string {
	graph, ok := getAdjList(words)
	if !ok {
		return "" //違反字典規則
	}
	var outBytes []byte
	visited := map[byte]bool{}
	cycleDetect := map[byte]bool{}
	for c, _ := range graph {
		r := dfs(graph, c, visited, cycleDetect, &outBytes)
		if !r {
			fmt.Println("detect cycle")
			return ""
		}
	}
	reverse(outBytes) //dfs直接append 所以需要再反轉
	return string(outBytes)
}

func getAdjList(words []string) (map[byte][]byte, bool) {
	adjList := map[byte][]byte{}
	for i, word := range words {
		nextword := ""
		if i+1 < len(words) {
			nextword = words[i+1]
		}

		isFindDiff := false
		for j := 0; j < len(word); j++ {
			b1 := word[j]
			if _, ok := adjList[b1]; !ok {
				adjList[b1] = []byte{}
			}

			if j < len(nextword) {
				b2 := nextword[j]
				if _, ok := adjList[b2]; !ok {
					adjList[b2] = []byte{}
				}

				if b1 != b2 && isFindDiff == false {
					isFindDiff = true
					adjList[b1] = append(adjList[b1], b2)
				}
			}
		}
		if isFindDiff == false && len(word) > len(nextword) && (i+1) < len(words) {//有相同前最，短的要在前
			fmt.Println("against dictionary rule")
			return adjList, false
		}

	}

	return adjList, true
}

//do topological sort
func dfs(graph map[byte][]byte, c byte, visited map[byte]bool, cycleDetect map[byte]bool, outBytes *[]byte) bool {
	if _,ok:=visited[c];ok{ //訪問過的不重複訪問印
		return true
	}

	if _,ok:=cycleDetect[c];ok{
		return false
	}

	cycleDetect[c]=true
	//loop dfs neighbor
	for _,neighbor:=range graph[c]{
		r:=dfs(graph,neighbor,visited,cycleDetect,outBytes)
		if !r{
			return false
		}
	}

	delete(cycleDetect,c)
	*outBytes=append(*outBytes,c)
	visited[c]=true
	return true
}

func reverse(outBytes []byte){
	//0 1 2 3
	//0 1 2 3 4

	for i,j:=0,len(outBytes)-1;i<len(outBytes)/2;i,j=i+1,j-1{
		outBytes[i],outBytes[j]=outBytes[j],outBytes[i]
	}
}
