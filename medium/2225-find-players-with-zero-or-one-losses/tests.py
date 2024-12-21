import unittest
from solution import Solution
from typing import List


class TestCase:
    def __init__(self, input: List[List[int]], want: List[List[int]]):
        self.want = want
        self.input = input


class TestSolution(unittest.TestCase):
    def test_equalSubstring(self):
        s = Solution()
        tests = [
            TestCase(
                [
                    [1, 3],
                    [2, 3],
                    [3, 6],
                    [5, 6],
                    [5, 7],
                    [4, 5],
                    [4, 8],
                    [4, 9],
                    [10, 4],
                    [10, 9],
                ],
                [[1, 2, 10], [4, 5, 7, 8]],
            ),
            TestCase([[2, 3], [1, 3], [5, 4], [6, 4]], [[1, 2, 5, 6], []]),
        ]

        for t in tests:
            s = Solution()
            self.assertEqual(t.want, s.findWinners(t.input))
