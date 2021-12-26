package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	t:=head
	r:=head
	for t!=nil&& r!=nil{
		t=t.Next
		r=r.Next
		if r==nil{
			return false
		}
		r=r.Next
		if t==r{
			return true
		}

	}
	return false
}