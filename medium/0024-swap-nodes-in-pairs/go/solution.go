package swap_nodes_in_pairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
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

func swapPairsValues(head *ListNode) *ListNode {
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
