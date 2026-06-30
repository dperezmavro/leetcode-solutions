package solution

import common "github.com/dperezmavro/leetcode-solutions/common/go"

func addTwoNumbers(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	sum := 0
	carry := 0
	res := &common.ListNode{}
	working_l := res
	for l1 != nil || l2 != nil {
		if carry == 1 {
			sum = 1
		}

		if l1 != nil {
			sum += l1.Val
		}

		if l2 != nil {
			sum += l2.Val
		}

		if sum >= 10 {
			sum = sum % 10
			carry = 1
		} else {
			carry = 0
		}

		working_l.Val = sum

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		if l1 == nil && l2 == nil {
			if carry == 1 {
				next := &common.ListNode{}
				working_l.Next = next
				working_l = next
			}
		} else {
			next := &common.ListNode{}
			working_l.Next = next
			working_l = next
		}

		sum = 0
	}

	if carry > 0 {
		working_l.Val = 1
		carry = 0
	} else if working_l.Val == 0 {
		working_l = nil
	}

	return res
}
