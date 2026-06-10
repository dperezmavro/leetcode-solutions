import unittest
from solution import Solution
from typing import List


class TestCase:
    def __init__(self, s1: str, s2: str, want: str):
        self.s1 = s1
        self.s2 = s2
        self.want = want


class TestSolution(unittest.TestCase):
    def test_equalSubstring(self):
        s = Solution()
        tests: List[TestCase] = [
            TestCase("ABCABC", "ABC", "ABC"),
            TestCase("ABABAB", "ABAB", "AB"),
            TestCase("LEET", "CODE", ""),
            TestCase("ABCDEF", "ABC", ""),
            TestCase("TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXX"),
        ]

        s = Solution()
        for t in tests:
            self.assertEqual(t.want, s.gcdOfStrings(t.s1, t.s2))
