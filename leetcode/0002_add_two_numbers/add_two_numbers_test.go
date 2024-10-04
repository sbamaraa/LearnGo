package add_two_numbers

import "testing"

/*
 * Return the given integer value as reversed ListNode.
 */
func intToListNode(a int) *ListNode {
	var root, prev *ListNode

	for {
		var temp *ListNode
		temp = &ListNode{
			Val: a % 10,
			Next: nil,
		}

		if prev != nil {
			prev.Next = temp
		} else {
			root = temp
		}

		prev = temp

		a /= 10

		if a == 0 {
			break;
		}
	}

	return root
}

func equals(l1 *ListNode, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	return l1 == nil && l2 == nil
}

func test(a int, b int, t *testing.T) {
	got := addTwoNumbers(intToListNode(a), intToListNode(b))
	wanted := intToListNode(a + b)

	if ! equals(got, wanted) {
		t.Error("FAIL :)")
	}
}

func TestCase1(t *testing.T) {
	test(342, 465, t)
}

func TestCase2(t *testing.T) {
	test(0, 0, t)
}

func TestCase3(t *testing.T) {
	test(9999999, 9999, t)
}
