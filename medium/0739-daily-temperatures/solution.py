from typing import List

class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        ans = [0] * len(temperatures)
        stack = []
        for idx in range(len(temperatures)):
            while stack and temperatures[stack[-1]] < temperatures[idx]:
                j = stack.pop()
                ans[j] = idx-j
            stack.append(idx)
        return ans
