package remove_duplicates_from_sorted_array

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	res := head
	for res != nil {
		if res.Next != nil && res.Val == res.Next.Val {
			res.Next = res.Next.Next
			continue
		}
		res = res.Next
	}

	return head
}
