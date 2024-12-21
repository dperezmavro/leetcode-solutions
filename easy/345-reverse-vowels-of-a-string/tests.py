import unittest
from solution import Solution
from typing import List


class TestCase:
    def __init__(self, s: str, want: str):
        self.s = s
        self.want = want


class TestSolution(unittest.TestCase):
    def test_equalSubstring(self):
        s = Solution()
        tests: List[TestCase] = [
            TestCase("IceCreAm", "AceCreIm"),
            TestCase("leetcode", "leotcede"),
        ]

        s = Solution()
        for t in tests:
            self.assertEqual(t.want, s.reverseVowels(t.s))
