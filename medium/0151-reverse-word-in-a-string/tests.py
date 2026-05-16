import unittest
from solution import Solution
from typing import List


class TestCase:
    def __init__(self, input: str, want: str):
        self.want = want
        self.input = input


class TestSolution(unittest.TestCase):
    def test_equalSubstring(self):
        tests: List[TestCase] = [
            TestCase("the sky is blue", "blue is sky the"),
            TestCase("  hello world  ", "world hello"),
            TestCase("a good   example", "example good a"),
        ]

        for t in tests:
            s = Solution()
            self.assertEqual(t.want, s.reverseWords(t.input))
