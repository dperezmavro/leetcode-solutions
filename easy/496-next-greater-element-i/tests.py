from typing import List
import unittest
from solution import Solution


class TestCase:
    def __init__(self, want: List[int], n1: List[int], n2: List[int]):
        self.want = want
        self.n1 = n1
        self.n2 = n2


class TestSolution(unittest.TestCase):
    def test_equalSubstring(self):
        s = Solution()
        tests = [
            TestCase([-1, 3, -1], [4, 1, 2], [1, 3, 4, 2]),
            TestCase([3, -1], [2, 4], [1, 2, 3, 4]),
            TestCase([7,7,7,7,7], [1,3,5,2,4], [6,5,4,3,2,1,7]),
        ]

        for t in tests:
            s = Solution()
            self.assertEqual(t.want, s.nextGreaterElement(t.n1, t.n2))
