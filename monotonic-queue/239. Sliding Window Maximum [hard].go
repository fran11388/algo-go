package main

import "fmt"


func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}
type Node struct {
	Prev *Node
	Next *Node
	Val  int
}

type Deque struct {
	Front *Node
	End   *Node
}

func initMonoticDeque(nums []int, i, j int) *Deque {
	deque:=&Deque{}
	for idx := i; idx <= j; idx++ {
		deque.monotonicPush(nums,idx)
	}
	return deque
}

func (d *Deque) getFirst() int {
	return d.Front.Val
}

func (d *Deque) slide(nums []int, i, j int) {
	if i > d.Front.Val { //pop front
		d.Front = d.Front.Next
		if d.Front != nil {
			d.Front.Prev = nil
		}
	}

	d.monotonicPush(nums, j)
}

func (d *Deque) monotonicPush(nums []int, j int) {
	newNum := nums[j]
	for true {
		node := d.End
		if node != nil {
			lastNum := nums[node.Val]
			if newNum > lastNum {
				d.popEnd()
			} else {
				break
			}
		} else {
			break
		}
	}

	d.pushEnd(j)
}

func (d *Deque) popEnd() *Node {
	if d.End != nil {
		if d.Front == d.End { //if only have one node
			d.Front, d.End = nil, nil
			return nil
		}

		tmp := d.End
		d.End = d.End.Prev
		if d.End != nil {
			d.End.Next = nil
		}

		return tmp
	}
	return nil
}

func (d *Deque) pushEnd(idx int) {
	node := &Node{Val: idx}
	if d.Front == nil {
		d.Front, d.End = node, node
		return
	}

	d.End.Next = node
	node.Prev = d.End
	d.End = node
}

func maxSlidingWindow(nums []int, k int) []int {
	result := []int{}
	deque := initMonoticDeque(nums, 0, k-1)
	for i, j := 0, k-1; j < len(nums); {
		result = append(result, nums[deque.getFirst()])
		i++
		j++
		if j >= len(nums) {
			break
		}
		deque.slide(nums, i, j)
	}
	return result
}


