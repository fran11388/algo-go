package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node==nil{
		return node
	}
	var queue []*Node
	queue = append(queue, node)
	cloneNodeMap:=make(map[int]*Node)
	//bfs
	visit:=make(map[int]bool)
	for len(queue)>0{
		popedNode:=queue[0]

		queue=queue[1:]

		v:=popedNode.Val
		var cloneNode *Node
		if cn,ok:=cloneNodeMap[v];ok{
			cloneNode=cn
		}else{
			cloneNode=&Node{Val: v}
			cloneNodeMap[v]=cloneNode
		}

		//clone node connect to neighbors
		for _,neighbor:=range popedNode.Neighbors{
			neighborVal:=neighbor.Val
			var newNeighbor *Node
			if cloneNeighbor,ok:=cloneNodeMap[neighborVal];ok{
				newNeighbor=cloneNeighbor
			}else{
				newNeighbor=&Node{Val: neighborVal}
				cloneNodeMap[neighborVal]=newNeighbor
			}

			cloneNode.Neighbors=append(cloneNode.Neighbors,newNeighbor)

			//if neighbor not explore, append to queue
			if   _,ok:=visit[neighborVal];!ok{
				queue=append(queue,neighbor)
				visit[neighborVal]=true //這邊就該標記visit了 因為是bfs  有可能在還沒pop出來之前多次append到queue
			}

		}
		visit[popedNode.Val]=true
	}

	return cloneNodeMap[node.Val]
}
