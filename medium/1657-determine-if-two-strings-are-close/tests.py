from typing import List
import unittest
from solution import Solution


class TestCase:
    def __init__(self, s1: str, s2: str, want: bool):
        self.s1 = s1
        self.s2 = s2
        self.want = want


class TestSolution(unittest.TestCase):
    def test_closeStrings(self):
        s = Solution()
        tests = [
            TestCase("abc", "bca", True),
            TestCase("a", "aa", False),
            TestCase("cabbba", "abbccc", True),
        ]

        for t in tests:
            s = Solution()
            self.assertEqual(t.want, s.closeStrings(t.s1, t.s2))
