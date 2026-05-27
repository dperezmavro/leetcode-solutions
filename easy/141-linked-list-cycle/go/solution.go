package linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	if head.Next == nil {
		return false
	}

	fast, slow := 0, 0
	fastNode := head
	slowNode := head

	for fastNode != nil && slowNode != nil {
		fast += 2
		slow += 1
		fastNode = fastNode.Next
		if fastNode == nil || fastNode.Next == nil {
			return false
		} else {
			fastNode = fastNode.Next
		}
		slowNode = slowNode.Next

		if fastNode == slowNode {
			return true
		}
	}

	return false
}
