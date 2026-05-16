class Solution:
    def reverseWords(self, s: str) -> str:
        parts = [x for x in s.split(" ") if len(x) > 0]
        return " ".join(parts[::-1])
