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
