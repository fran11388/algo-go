package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	//find meet point
	t := head
	r := head
	var meetPoint *ListNode
	for t != nil && r != nil {
		t = t.Next
		r = r.Next
		if r == nil {
			return nil
		}
		r = r.Next
		if t == r {
			//fine the meet point
			meetPoint=t
			break
		}
	}
	if meetPoint==nil{
		return nil
	}

	//move rabbit to head, and move tuttle and rabbit equal spped
	r = head
	for true {
		if r == t { //注意這邊應該先判斷
			return t
		}
		r = r.Next
		t = t.Next

	}

	return nil
}
