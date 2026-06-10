import unittest
from solution import Solution
from typing import List


class TestCase:
    def __init__(self, input: List[int], want: int):
        self.input = input
        self.want = want


class TestSolution(unittest.TestCase):
    def test_equalSubstring(self):
        s = Solution()
        tests: List[TestCase] = [
            TestCase([-5, 1, 5, 0, -7], 1),
            TestCase([-4, -3, -2, -1, 4, 3, 2], 0),
        ]

        s = Solution()
        for t in tests:
            self.assertEqual(t.want, s.largestAltitude(t.input))
