import unittest
from solution import Solution


class TestCase:
    def __init__(self, want: int, s: str, t: str, mC: int):
        self.want = want
        self.s = s
        self.t = t
        self.mC = mC


class TestSolution(unittest.TestCase):
    def test_equalSubstring(self):
        s = Solution()
        tests = [
            TestCase(3, "abcd", "bcdf", 3),
            TestCase(1, "abcd", "cdef", 3),
            TestCase(1, "abcd", "acde", 0),
            TestCase(2, "krrgw", "zjxss", 19),
            TestCase(3, "baaa", "caaa", 0),
        ]

        for t in tests:
            s = Solution()
            self.assertEqual(t.want, s.equalSubstring(t.s, t.t, t.mC))
