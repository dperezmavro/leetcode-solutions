from typing import List
from itertools import combinations


class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        if digits == "":
            return []

        mappings = {
            "2": ["a", "b", "c"],
            "3": ["d", "e", "f"],
            "4": ["g", "h", "i"],
            "5": ["j", "k", "l"],
            "6": ["m", "n", "o"],
            "7": ["p", "q", "r", "s"],
            "8": ["t", "u", "v"],
            "9": ["w", "x", "y", "z"],
        }
        results = mappings[digits[0]]

        for d in digits[1:]:
            results = [l + x for l in results for x in mappings[d]]
        return results


if __name__ == "__main__":
    s = Solution()
    print(s.letterCombinations("23"))
    print(s.letterCombinations(""))
    print(s.letterCombinations("2"))
    print(s.letterCombinations("234"))
