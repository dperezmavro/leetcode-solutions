from typing import List
import unittest
from solution import Solution


class TestCase:
    def __init__(self, want: List[int], n: List[int]):
        self.want = want
        self.n = n


class TestSolution(unittest.TestCase):
    def test_equalSubstring(self):
        s = Solution()
        tests = [
            TestCase([1, 1, 4, 2, 1, 1, 0, 0], [73, 74, 75, 71, 69, 72, 76, 73]),
            TestCase([1, 1, 1, 0], [30, 40, 50, 60]),
            TestCase([1, 1, 0], [30, 60, 90]),
        ]

        for t in tests:
            s = Solution()
            self.assertEqual(t.want, s.dailyTemperatures(t.n))
