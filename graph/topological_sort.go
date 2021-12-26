package main

import "fmt"

func main(){
	g:=map[string][]string{
		"a":[]string{"b","e"},
		"b":[]string{"c"},
		"c":[]string{"d"},
		"d":[]string{},
		"e":[]string{"f"},
		"f":[]string{"d"},
	}
	//g:=map[string][]string{
	//	"a":[]string{"b"},
	//	"b":[]string{"a"},
	//
	//}

	fmt.Println(topologicalSort(g))
}

func topologicalSort(graph map[string][]string)[]string{
	result:=make([]string,len(graph))
	insertPos:=len(graph)-1
	visit:=make(map[string]bool)
	for node,_:=range graph{ //dfs loop through every node
		dfs(graph,node,&insertPos,visit,result)

	}
	return result
}

func dfs(graph map[string][]string,node string,insertPos *int,visit map[string]bool,result []string){

	if _,ok:=visit[node];ok{
		return
	}
	visit[node]=true
	for _,neighbor:=range graph[node]{
		dfs(graph,neighbor,insertPos,visit,result)
	}

	result[*insertPos]=node
	*insertPos=(*insertPos)-1

}
