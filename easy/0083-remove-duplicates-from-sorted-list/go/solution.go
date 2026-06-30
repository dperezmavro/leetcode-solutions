package remove_duplicates_from_sorted_array

import common "github.com/dperezmavro/leetcode-solutions/common/go"

func deleteDuplicates(head *common.ListNode) *common.ListNode {
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
