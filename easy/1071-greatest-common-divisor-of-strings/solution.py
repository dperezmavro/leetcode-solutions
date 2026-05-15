class Solution:
    def gcdOfStrings(self, str1: str, str2: str) -> str:
        chars_1 = set(str1)
        chars_2 = set(str2)
        for c in chars_2:
            if c not in chars_1:
                return ""
        if len(chars_1) != len(chars_2):
            return ""
        if len(str2) > len(str1):
            str2, str1 = str1, str2

        # done cleaning, now to compute

        while str1 and str2:
            str1 = []
