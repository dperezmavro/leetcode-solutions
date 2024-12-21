class Solution:
    def reverseVowels(self, s: str) -> str:
        all_vowels = set(["I", "i", "a", "A", "e", "E", "o", "O", "u", "U"])
        res = [x for x in s]
        vowels = [x for x in s if x in all_vowels]
        for idx, c in enumerate(res):
            if c in all_vowels:
                res[idx] = vowels.pop()
        return "".join(res)
