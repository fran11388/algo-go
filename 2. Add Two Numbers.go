package main

func main() {

}

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var prevNode *ListNode
	var rootNode *ListNode

	carry := 0
	for {
		if l1 == nil && l2 == nil {
			break
		}
		v := carry
		if l1 != nil {
			v += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v += l2.Val
			l2 = l2.Next
		}

		node := &ListNode{}
		node.Val = v % 10
		carry = v / 10

		if prevNode != nil {
			prevNode.Next = node

		} else {
			rootNode = node

		}
		prevNode = node
	}
	if carry>0{
		node:=&ListNode{}
		node.Val=carry
		prevNode.Next=node

	}

	return rootNode
}
