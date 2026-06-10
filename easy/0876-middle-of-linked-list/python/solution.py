from typing import Optional

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def middleNode(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head:
            return None

        fast_index = 0
        slow_index = 0
        fast_node = head
        slow_node = head
        while fast_node.next:
            if (fast_index // 2) + 1 > slow_index:
                slow_index += 1
                slow_node = slow_node.next
            fast_node = fast_node.next
            fast_index += 1

        return slow_node

