from typing import Optional


# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def removeNthFromEnd(
        self,
        head: Optional[ListNode],
        n: int,
    ) -> Optional[ListNode]:

        if not head:
            return head

        fast = head
        slow = head

        for _ in range(n):
            fast = fast.next

        while fast and fast.next:
            fast = fast.next
            slow = slow.next

        # condition where n == len(ListNode)
        if not fast:
            return head.next

        if slow and slow.next:
            slow.next = slow.next.next
        else:
            return None

        return head
