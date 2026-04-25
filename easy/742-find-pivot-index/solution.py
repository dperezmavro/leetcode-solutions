from typing import List


class Solution:
    def pivotIndex(self, nums: List[int]) -> int:
        answer = -1
        sumTotal = 0
        partialSums: List[int] = [0] * len(nums)

        for idx, v in enumerate(nums):
            sumTotal += v
            partialSums[idx] = sumTotal

        for idx, v in enumerate(nums):
            leftSum = partialSums[idx] - v
            rightSum = sumTotal - partialSums[idx]
            if rightSum == leftSum:
                return idx
        return answer


s = Solution()
print(s.pivotIndex([1, 7, 3, 6, 5, 6]))
print(s.pivotIndex([1, 2, 3]))
print(s.pivotIndex([2, 1, -1]))
