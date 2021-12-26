package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//time:O(n)
//space:O(1)
func isPalindrome(head *ListNode) bool {
	//find moddle point
	r := head
	t := head
	for r != nil && t != nil {
		r = r.Next
		t = t.Next
		if r == nil {
			break
		}
		r = r.Next
		if r == nil {
			break
		}
	}
	//t at  middle node now

	reversedRoot := reverseLinklist(t)
	r = head
	for reversedRoot != nil {
		if reversedRoot.Val != r.Val {
			return false
		}
		reversedRoot = reversedRoot.Next
		r = r.Next
	}
	return true

}

func reverseLinklist(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	var root *ListNode
	for curr != nil {
		root = curr
		tmpNext := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmpNext
	}

	return root
}
