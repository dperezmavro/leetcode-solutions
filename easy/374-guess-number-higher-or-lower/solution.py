# The guess API is already defined for you.
# @param num, your guess
# @return -1 if num is higher than the picked number
#          1 if num is lower than the picked number
#          otherwise return 0
# def guess(num: int) -> int:
from math import floor


def guessFunc(n: int):
    k = n

    def guess(n2: int) -> int:
        if n2 == k:
            return 0
        elif n2 > k:
            return -1
        else:
            return 1

    return guess


class Solution:
    def guessNumber(self, n: int, k: int) -> int:
        if n == 1:
            return 1
        min_range = 0
        max_range = n
        res = -1
        guess = guessFunc(k)
        while res != 0:
            g = min_range + int((max_range - min_range) / 2)
            res = guess(g)
            print(min_range, max_range, g, res)
            if res == 0:
                return g
            elif res == 1:
                min_range = g + 1
            else:
                max_range = g - 1
        return -1
