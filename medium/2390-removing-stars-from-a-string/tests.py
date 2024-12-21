import unittest
from solution import Solution
from typing import List


class TestCase:
    def __init__(self, input: str, want: str):
        self.want = want
        self.input = input


class TestSolution(unittest.TestCase):
    def test_equalSubstring(self):
        tests = [
            TestCase("leet**cod*e", "lecoe"),
            TestCase("erase*****", ""),
        ]

        s = Solution()
        for tt in tests:
            self.assertEqual(tt.want, s.removeStars(tt.input))
