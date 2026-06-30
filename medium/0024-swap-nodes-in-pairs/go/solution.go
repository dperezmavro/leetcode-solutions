package swap_nodes_in_pairs

import common "github.com/dperezmavro/leetcode-solutions/common/go"

func swapPairs(head *common.ListNode) *common.ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	dummy := &common.ListNode{Next: head}
	prev := dummy

	for prev.Next != nil && prev.Next.Next != nil {
		first := prev.Next
		second := first.Next

		first.Next = second.Next
		second.Next = first
		prev.Next = second

		prev = first
	}

	return dummy.Next

}

func swapPairsValues(head *common.ListNode) *common.ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	curr := head
	tmp := 0
	for curr != nil && curr.Next != nil {
		tmp = curr.Next.Val
		curr.Next.Val = curr.Val
		curr.Val = tmp
		curr = curr.Next.Next
	}

	return head
}
