package main

func main(){

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
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//declare a dummy head, curr
	//loop compare l1,l2 , and move curr
	if l1==nil{
		return l2
	}
	if l2==nil{
		return l1
	}

	var curr *ListNode
	var dummyHead *ListNode
	if l1.Val<l2.Val{
		curr=l1
		dummyHead=l1
		l1=l1.Next
	}else{
		curr=l2
		dummyHead=l2
		l2=l2.Next
	}

	for l1!=nil&&l2!=nil{
		if l1.Val<l2.Val{
			curr.Next=l1
			l1=l1.Next
		}else{
			curr.Next=l2
			l2=l2.Next
		}
		curr=curr.Next

	}

	if l1==nil{
		curr.Next=l2
	}else{
		curr.Next=l1
	}

	return dummyHead
}