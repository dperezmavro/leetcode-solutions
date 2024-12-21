import unittest
from solution import Solution
from typing import List


class TestCase:
    def __init__(self, n: int, pick: int):
        self.n = n
        self.pick = pick


class TestSolution(unittest.TestCase):
    def test_equalSubstring(self):
        s = Solution()
        tests: List[TestCase] = [
            TestCase(10, 6),
            TestCase(1, 1),
            TestCase(2, 1),
            TestCase(2, 2),
        ]

        for t in tests:
            s = Solution()
            self.assertEqual(t.pick, s.guessNumber(t.n, t.pick))
