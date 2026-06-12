import unittest
from solution import Solution


class TestCase:
    def __init__(self, i: str, w: int):
        self.input = i
        self.want = w


class TestSolution(unittest.TestCase):
    def test_solution(self):
        tests = [
            TestCase(i="abcabcbb", w=3),
            TestCase(i="bbbbb", w=1),
            TestCase(i="pwwkew", w=3),
            TestCase(i="tmmzuxt", w=5),
        ]
        for t in tests:
            s = Solution()
            r = s.lengthOfLongestSubstring(t.input)
            self.assertEqual(t.want, r)


if __name__ == "__main__":
    unittest.main()
