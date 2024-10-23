package swap_nodes

type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head;
	}

	fHead := head.Next
	head.Next = swapPairs(head.Next.Next)
	fHead.Next = head

	return fHead
}
