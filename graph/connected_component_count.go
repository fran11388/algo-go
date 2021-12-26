package main

import "fmt"

func main() {
	graph := map[string][]string{
		"a": []string{"b"},
		"b": []string{"a"},
		"c": []string{"d"},
		"d": []string{"c", "e"},
		"e": []string{"d"},
		"f": []string{"g"},
		"g": []string{"f", "h"},
		"h": []string{"g"},
	}

	fmt.Println(connectedComponentCount(graph))
}

func connectedComponentCount(graph map[string][]string) int {
	count := 0
	visit := make(map[string]bool)
	for node, _ := range graph {
		if exploreByDfs(graph, node, visit) {
			count++
		}
	}
	return count
}

func exploreByDfs(graph map[string][]string, node string, visit map[string]bool) bool{
	if _,ok:=visit[node];ok{
		return  false
	}
	visit[node]=true

	for _,neighbor:=range graph[node]{
		exploreByDfs(graph,neighbor,visit)
	}

	return true

}
