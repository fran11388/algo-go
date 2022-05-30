package main

import "fmt"

func main() {
	data := getSimpleAdjList()
	printByDFS(data)
}

func printByDFS(graph map[int][]int) {
	visit := map[int]bool{}
	for k, _ := range graph {
		visit[k] = false
	}

	orderlist := []int{}
	for node, _ := range graph {
		topologicalDFS(visit, graph, node, &orderlist)
	}

	reverseInt(orderlist)
	fmt.Println(orderlist)
}

func topologicalDFS(visit map[int]bool, graph map[int][]int, node int, orderlist *[]int) {
	if visit[node] {
		return
	}

	for _, neighbor := range graph[node] {
		topologicalDFS(visit, graph, neighbor, orderlist)
	}

	*orderlist = append(*orderlist, node)
	visit[node] = true

}

func reverseInt(orderlist []int) {
	j := len(orderlist) - 1
	for i := 0; i < len(orderlist)/2; i++ {
		orderlist[i], orderlist[j] = orderlist[j], orderlist[i]
		j--
	}
}

func getSimpleAdjList() map[int][]int {
	m := map[int][]int{}
	m[0] = []int{1, 2}
	m[1] = []int{3}
	m[2] = []int{3}
	m[3] = []int{4, 5}
	m[4] = []int{6}
	m[5] = []int{6}
	m[6] = []int{}
	return m
}
