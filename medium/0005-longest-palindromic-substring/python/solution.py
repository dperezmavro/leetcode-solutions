from typing import Tuple


class Solution:
    def longestPalindrome(self, s: str) -> str:
        res = ""

        if len(s) <= 1:
            return s

        for i, c in enumerate(s):
            max_o = self.expand_centre(s, i, i)
            max_e = self.expand_centre(s, i, i + 1)
            str_o = s[max_o[0] : max_o[1] - 1]
            str_e = s[max_e[0]+1 : max_e[1]]

            # print(f"res {res} str_e {str_e} str_o {str_o} max_o {max_o} max_e {max_e}")
            if len(str_o) > len(res):
                res = str_o
            if len(str_e) > len(res):
                res = str_e
        return res

    def expand_centre(self, s: str, left: int, right: int) -> Tuple[int]:
        # print(f"scaning {s} around {left} {right}")
        while left >= 0 and right < len(s) and s[left] == s[right]:
            # print(f"matching on {left} {right} {s[left:right]}")
            left -= 1
            right += 1
        return (left, right)
