class Solution:
    def reverseOnlyLetters(self, s: str) -> str:
        res = [x for x in s]
        alphas = [x for x in s if self.isLetter(x)]
        alphas = alphas[::-1]
        offset = 0
        for idx, v in enumerate(alphas):
            while not self.isLetter(res[idx + offset]):
                offset += 1
            res[idx + offset] = v

        return "".join(res)

    def isLetter(self, a: str) -> bool:
        return (a >= "a" and a <= "z") or (a >= "A" and a <= "Z")


s = Solution()
print(s.reverseOnlyLetters("ab-cd"))
print(s.reverseOnlyLetters("a-bC-dEf-ghIj"))
print(s.reverseOnlyLetters("Test1ng-Leet=code-Q!"))
