package add_two_numbers

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode{
	var root, prev *ListNode
	var r = 0

	for l1 != nil || l2 != nil || r != 0 {
		if l1 != nil {
			r += l1.Val
		}

		if l2 != nil {
			r += l2.Val
		}

		// Create the node
		temp := &ListNode {
			Val: r % 10,
			Next: nil,
		}

		// Connection to the current node from prev node
		if prev != nil {
			prev.Next = temp
		} else {
			root = temp
		}

		prev = temp

		r = r / 10

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	return root
}

