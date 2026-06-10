import unittest
from typing import List
from solution import Solution, ListNode


class TestCase:
    def __init__(self, list: List[int], hasCycle=bool):
        head = ListNode(x=list[0])
        node = head
        for i in list[1:]:
            nn = ListNode(x=i)
            node.next = nn
            node = nn
        self.head = head
        self.hasCycle = hasCycle

        if self.hasCycle:
            node.next = head.next


class TestSolution(unittest.TestCase):
    def test_Solution(self):
        tests = [
            TestCase(list=[1], hasCycle=False),
            TestCase(list=[1, 2], hasCycle=True),
            TestCase(list=[1, 2, 3, 4, 5, 6, 7], hasCycle=True),
        ]
        for t in tests:
            s = Solution()
            self.assertEqual(t.hasCycle, s.hasCycle(t.head))
