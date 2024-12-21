class Solution:
    def __init__(self):
        self.vowels = set({"a", "e", "i", "o", "u"})

    def maxVowels(self, s: str, k: int) -> int:
        maxVowels = 0
        vowelsInWindow = 0
        for idx, c in enumerate(s):
            if c in self.vowels:
                vowelsInWindow += 1
            if idx >= k:
                if s[idx - k] in self.vowels:
                    vowelsInWindow -= 1

            maxVowels = max(maxVowels, vowelsInWindow)

        return maxVowels
