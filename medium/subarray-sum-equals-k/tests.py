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
            TestCase(2, [1, 1, 1], 2),
            TestCase(2, [1, 2, 3], 3),
            TestCase(1, [-1, -1, 1], 0),
            TestCase(3, [1, -1, 0], 0),
            TestCase(1, [100, 1, 2, 3, 4], 6),
            
            TestCase(1, [28, 54, 7, -70, 22, 65, -6], 100),
        ]

        for t in tests:
            s = Solution()
            self.assertEqual(t.want, s.subarraySum(t.n, t.k))
