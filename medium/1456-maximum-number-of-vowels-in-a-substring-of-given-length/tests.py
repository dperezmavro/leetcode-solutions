import unittest
from solution import Solution


class TestCase:
    def __init__(self, want: int, s: str, k: int):
        self.want = want
        self.s = s
        self.k = k


class TestSolution(unittest.TestCase):
    def test_equalSubstring(self):
        s = Solution()
        tests = [
            TestCase(3, "abciiidef", 3),
            TestCase(2, "aeiou", 2),
            TestCase(2, "leetcode", 3),
        ]

        for t in tests:
            s = Solution()
            self.assertEqual(t.want, s.maxVowels(t.s, t.k))
