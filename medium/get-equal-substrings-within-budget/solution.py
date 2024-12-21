class Solution:
    def equalSubstring(self, s: str, t: str, maxCost: int) -> int:
        currentCost = length = left = 0

        for right, c in enumerate(s):
            currentCost += abs(ord(c) - ord(t[right]))
            while currentCost > maxCost:
                currentCost -= abs(ord(s[left]) - ord(t[left]))
                left += 1
                length = max(length, right + 1 - left)
            length = max(length, right + 1 - left)

        return length
