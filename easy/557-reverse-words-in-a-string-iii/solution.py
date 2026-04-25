class Solution:
    def reverseWords(self, s: str) -> str:
        partials = s.split(" ")
        for idx, v in enumerate(partials):
            partials[idx] = v[::-1]
        return ' '.join(partials)


s = Solution()
print(s.reverseWords("Let's take LeetCode contest"))
print(s.reverseWords("Mr Ding"))
