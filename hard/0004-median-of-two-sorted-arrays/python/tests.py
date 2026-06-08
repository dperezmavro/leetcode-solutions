import unittest
from solution import Solution
from typing import List


class TestCase:
    def __init__(self, input_1: List[int], input_2: List[int], want: int):
        self.want = want
        self.input_1 = input_1
        self.input_2 = input_2


class TestSolution(unittest.TestCase):
    def test_equalSubstring(self):
        def assess(self, t: TestCase):
            s = Solution()
            m = s.findMedianSortedArraysSlow(t.input_1, t.input_2)
            self.assertEqual(t.want, m)

        tests = [TestCase([1, 3], [2], 2), TestCase([1, 2], [3, 4], 2.5)]

        for t in tests:
            assess(self, t)
