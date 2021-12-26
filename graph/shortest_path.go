package main

import "fmt"

func main() {
	g := map[string][]string{
		"a": []string{"b", "d"},
		"b": []string{"a", "c"},
		"c": []string{"b", "e"},
		"d": []string{"a", "e"},
		"e": []string{"d", "c"},
	}

	fmt.Println(shortestPath(g, "a", "c"))
}

type node struct{
	Name string
	Distance int
}

func shortestPath(graph map[string][]string, nodeA, nodeB string) int {
	var queue []node
	queue = append(queue, node{Name: nodeA,Distance: 0})
	visit:=make(map[string]bool)
	//visit[nodeA]=true
	for len(queue)>0{
		//pop
		n :=queue[0]
		queue=queue[1:]
		visit[n.Name]=true

		if n.Name==nodeB{
			return n.Distance
		}

		for _,neighbor:=range graph[n.Name]{
			if _,ok:=visit[neighbor];!ok{
				queue=append(queue, node{Name: neighbor,Distance: n.Distance+1})
			}

		}

	}
	return -1
}
