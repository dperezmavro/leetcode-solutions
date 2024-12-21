from typing import List
from collections import defaultdict


class Solution:
    def subarraySum(self, nums: List[int], k: int) -> int:
        answer = 0
        curr = 0
        sums = defaultdict(int)
        sums[0] = 1
        
        for n in nums:
            curr += n
            answer += sums[curr - k]
            sums[curr] += 1

        return answer

    def subarraySumOLD(self, nums: List[int], k: int) -> int:
        left = answer = 0
        currSum = 0
        for right, v in enumerate(nums):
            print(
                "1. currSum {} l {} r {} nums {}".format(
                    currSum, left, right, nums[left : right + 1]
                )
            )
            currSum += v
            if currSum == k:
                answer += 1

            while left < right and (currSum > k or right == len(nums) - 1):
                print(
                    "2. currSum {} l {} r {} nums {}".format(
                        currSum, left, right, nums[left : right + 1]
                    )
                )
                currSum -= nums[left]
                left += 1
                if currSum == k:
                    answer += 1
                print(
                    "3. currSum {} l {} r {} nums {}".format(
                        currSum, left, right, nums[left : right + 1]
                    )
                )

        return answer
