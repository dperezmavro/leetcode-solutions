from typing import List
import unittest
from solution import Solution


class TestCase:
    def __init__(self, want: int, n: List[int], k: int):
        self.want = want
        self.n = n
        self.k = k


class TestSolution(unittest.TestCase):
    def test_equalSubstring(self):
        s = Solution()
        tests = [
            TestCase(2, [1, 1, 2, 1, 1], 3),
            TestCase(0, [2, 4, 6], 1),
            TestCase(16, [2, 2, 2, 1, 2, 2, 1, 2, 2, 2], 2),
        ]

        for t in tests:
            s = Solution()
            self.assertEqual(t.want, s.numberOfSubarrays(t.n, t.k))
