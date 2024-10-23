package swap_nodes

import (
	"reflect"
	"testing"
)

func linkedListToSlice(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func sliceToLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil;
	}

	head := &ListNode{Val: nums[0]}
	current := head
	for _, val := range nums[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

func TestCase1(t *testing.T) {
	got := linkedListToSlice(swapPairs(sliceToLinkedList([]int{1, 2, 3, 4})))
	wanted := []int{2, 1, 4, 3}

	if ! reflect.DeepEqual(got, wanted) {
		t.Errorf("got: %v, wanted: %v", got, wanted)
	}
}

func TestCase2(t *testing.T) {
	got := linkedListToSlice(swapPairs(sliceToLinkedList([]int{})))
	wanted := []int{}

	if ! reflect.DeepEqual(got, wanted) {
		t.Errorf("got: %v, wanted: %v", got, wanted)
	}
}

func TestCase3(t *testing.T) {
	got := linkedListToSlice(swapPairs(sliceToLinkedList([]int{1})))
	wanted := []int{1}

	if ! reflect.DeepEqual(got, wanted) {
		t.Errorf("got: %v, wanted: %v", got, wanted)
	}
}

func TestCase4(t *testing.T) {
	got := linkedListToSlice(swapPairs(sliceToLinkedList([]int{1, 2, 3})))
	wanted := []int{2, 1, 3}

	if ! reflect.DeepEqual(got, wanted) {
		t.Errorf("got: %v, wanted: %v", got, wanted)
	}
}

func TestCase5(t *testing.T) {
	got := linkedListToSlice(swapPairs(sliceToLinkedList([]int{1, 2, 3, 4, 5, 6})))
	wanted := []int{2, 1, 4, 3, 6, 5}

	if ! reflect.DeepEqual(got, wanted) {
		t.Errorf("got: %v, wanted: %v", got, wanted)
	}
}
