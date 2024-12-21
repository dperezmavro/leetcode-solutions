from typing import List


class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        left = currSum = totalSum = 0
        minSize = len(nums)
        for right, v in enumerate(nums):
            currSum += v
            totalSum += v
            while currSum >= target:
                currSum -= nums[left]
                minSize = min(minSize, right - left + 1)
                left += 1
        if totalSum < target:
            return 0
        return minSize


s = Solution()
print(2, s.minSubArrayLen(7, [2, 3, 1, 2, 4, 3]))
print(1, s.minSubArrayLen(4, [1, 4, 4]))
print(0, s.minSubArrayLen(11, [1, 1, 1, 1, 1, 1, 1, 1]))
