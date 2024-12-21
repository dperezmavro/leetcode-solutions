class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        c = 0
        res = []
        while c < len(word1) or c < len(word2):
            if c < len(word1):
                res.append(word1[c])
            if c < len(word2):
                res.append(word2[c])

            c+=1
        return "".join(res)
        
