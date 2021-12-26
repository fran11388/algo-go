package main

import "fmt"

func main(){
	preq:=[][]int{{0,1},{1,0}}
	n:=2
	fmt.Println(canFinish(n,preq))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	//init graph
	graph:=map[int][]int{}
	for _,prerequisite:=range prerequisites{
		graph[prerequisite[0]]=append(graph[prerequisite[0]],prerequisite[1])
	}

	canFinishCourse:=map[int]bool{} //這行做cache很重要  不然performance test不會過
	visit:=make(map[int]bool)
	//dfs every node in graph
	for n,_:=range graph{
		r:=dfs(graph,n,visit,canFinishCourse)
		if !r{
			return false
		}
	}

	return true
}

func dfs(graph map[int][]int, n int, visit map[int]bool,canFinishCourse map[int]bool)bool{
	if v,ok:=canFinishCourse[n];ok{
		return v
	}

	//dfs every neighbor
	if _,ok:=visit[n];ok{
		return false //detect cycle
	}
	visit[n]=true
	for _,neighbor:=range graph[n]{
		r:=dfs(graph,neighbor,visit,canFinishCourse)
		if !r{
			canFinishCourse[n]=false
			return false
		}
	}
	delete(visit,n)
	canFinishCourse[n]=true
	return true
}

