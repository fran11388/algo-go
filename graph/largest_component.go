package main

import "fmt"

func main(){
	g:=map[string][]string{
		"a":[]string{"b"},
		"b":[]string{"a"},
		"c":[]string{"d"},
		"d":[]string{"c","e"},
		"e":[]string{"d"},
		"f":[]string{"g"},
		"g":[]string{"f","h"},
		"h":[]string{"g","i"},
		"i":[]string{"h"},
	}

	fmt.Println(largestComponent(g))// should be 4
}

func largestComponent(graph map[string][]string)int{
	largest:=0
	//iterate every node
	visit:=make(map[string]bool)
	for node,_:=range graph{
		e:=exploreSize(graph,node,visit)
		if e>largest{
			largest=e
		}
	}
	return largest
}

func exploreSize(graph map[string][]string,node string ,visit map[string]bool)int{
	if _,ok:=visit[node];ok{
		//ignore visited
		return 0
	}
	visit[node]=true
	count:=1
	for _,neighbot:=range graph[node]{
		count+=exploreSize(graph,neighbot,visit)
	}
	return count

}