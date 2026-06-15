from typing import Tuple


class Solution:
    def longestPalindrome(self, s: str) -> str:
        if len(s) <= 1:
            return s

        start, end = 0, 0
        for i in range(len(s)):
            # have to treat two different cases
            # to keep it simple
            # odd palindromes
            pal_odd = self.expand_centre(s, i, i)
            # even palindromes
            pal_even = self.expand_centre(s, i, i + 1)
            if pal_odd[1] - pal_odd[0] > end - start:
                start, end = pal_odd
            if pal_even[1] - pal_odd[1] > end - start:
                start, end = pal_even
        return s[start : end + 1]

    def expand_centre(self, s: str, left: int, right: int) -> Tuple[int, int]:
        while left >= 0 and right < len(s) and s[left] == s[right]:
            left -= 1
            right += 1
        return (left + 1, right - 1)
