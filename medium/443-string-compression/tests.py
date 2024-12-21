import unittest
from solution import Solution
from typing import List


class TestCase:
    def __init__(self, input: List[str], want: int):
        self.input = input
        self.want = want


class TestSolution(unittest.TestCase):
    def test_equalSubstring(self):
        tests: List[TestCase] = [
            TestCase(["a", "a", "b", "b", "c", "c", "c"], 6),
            TestCase(["a"], 1),
            TestCase(
                ["a", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b"], 4
            ),
        ]

        for t in tests:
            s = Solution()
            self.assertEqual(t.want, s.compress(t.input))
