import unittest
from solution import Solution, ListNode
from typing import List


class TestCase:
    def __init__(self, input: List[int], want: int):
        self.want = want
        self.input = input

        head = ListNode(val=input[0])
        node = head
        for i in input[1:]:
            node.next = ListNode(val=i)
            node = node.next
        self.head = head


class TestSolution(unittest.TestCase):
    def test_middleNode(self):
        tests = [
            TestCase(input=[1, 2, 3, 4, 5], want=3),
            TestCase(input=[1, 2, 3, 4, 5, 6], want=4),
        ]

        for t in tests:
            s = Solution()
            res = s.middleNode(head=t.head)
            self.assertEqual(t.want, res.val)
