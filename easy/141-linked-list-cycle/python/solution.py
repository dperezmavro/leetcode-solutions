from typing import Optional


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        fast = head
        slow = head
        index = 0
        while fast and fast.next:
            if fast == slow:
                if index > 0:
                    return True
                else:
                    index = 1
            slow = slow.next
            fast = fast.next.next
        return False
