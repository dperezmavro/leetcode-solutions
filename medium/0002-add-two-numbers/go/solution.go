package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := 0
	carry := 0
	res := &ListNode{}
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
				next := &ListNode{}
				working_l.Next = next
				working_l = next
			}
		} else {
			next := &ListNode{}
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
