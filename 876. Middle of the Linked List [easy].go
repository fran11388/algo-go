package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	r:=head
	t:=head
	if head.Next==nil{
		return head
	}
	for true{
		r=r.Next
		t=t.Next
		if r.Next==nil{
			return t
		}
		r=r.Next
	}
	return nil
}