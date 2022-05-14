package union_find

func countComponents(n int, edges [][]int) int {
	arr := make([]int, n, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}

	size := make([]int, n, n) //用來紀錄每個component大小 讓union的時候能夠選擇較小的component其parent設為較大的component
	for i := 0; i < n; i++ {
		size[i] = 1
	}

	count := n
	for i := 0; i < len(edges); i++ {
		edge := edges[i]
		v1, v2 := edge[0], edge[1]
		count -= combine(arr, size, v1, v2)
	}
	return count
}

func combine(arr []int, size []int, v1 int, v2 int) int {
	v1 = find(arr, v1)
	v2 = find(arr, v2)

	if v1 == v2 {
		return 0 //代表已經在同個component
	}

	if size[v1] > size[v2] {
		size[v1] += size[v2]
		arr[v2] = v1
	} else {
		size[v2] += size[v1]
		arr[v1] = v2
	}

	return 1
}

func find(arr []int, v int) int {
	if arr[v] == v {

		return v
	}

	arr[v] = find(arr, arr[v]) //找到root時同時也更新 加速之後查找(path compression)
	return arr[v]
}
