package main

import "fmt"

func main() {
	//g := map[string][]string{
	//	"f":[]string{"g","i"},
	//	"g":[]string{"h"},
	//	"h":[]string{},
	//	"i":[]string{"g","k"},
	//	"j":[]string{"i"},
	//	"k":[]string{},
	//}
	g := map[string][]string{
		"a":[]string{"b","d"},
		"b":[]string{"c"},
		"c":[]string{"a"},

	}
	fmt.Println(hasPathDfs(g,"a","d",make(map[string]bool)))


}

func hasPathDfs(graph map[string][]string, src string, dst string,visit map[string]bool) bool {

	fmt.Println(src)
	if src==dst{
		fmt.Println("src==dst")
		return true
	}
	if _,ok:=visit[src];ok{
		fmt.Println("visited")
		return false
	}

	visit[src]=true
	for _,neighbor:=range graph[src]{
		if hasPathDfs(graph,neighbor,dst,visit){
			return true
		}
	}

	return false
}

func hasPathBfs(graph map[string][]string, src string, dst string) bool {
	queue:=[]string{}
	queue=append(queue,src)

	for len(queue)>0{
		node:=queue[0]
		queue=queue[1:]
		if node==dst{
			return true
		}
		queue=append(queue,graph[node]...)
	}

	return false
}
