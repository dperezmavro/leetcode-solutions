import unittest
from solution import Solution


class TestCase():
    def __init__(self, input: str, want: str):
        self.input = input
        self.want = want

class TestSolution(unittest.TestCase):
    def test_solution(self):
        tests = [
            # TestCase(input="babad", want="bab"),
            TestCase(input="cbbd", want="bb"),
        ]

        for t in tests:
            s = Solution()
            res = s.longestPalindrome(t.input)
            self.assertEqual(t.want, res)

if __name__ == "__main__":
    unittest.main()