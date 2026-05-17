from typing import List


class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        nums = [x for x in nums2]
        for i in nums1:
            nums.append(i)
        nums.sort()
        if len(nums) % 2 == 0:
            midpoint = int((len(nums) - 1) / 2)
            numA = nums[midpoint : midpoint + 2]
            return sum(numA) / 2
        else:
            return float(nums[int((len(nums) - 1) / 2)])


s = Solution()
print(s.findMedianSortedArrays([1, 3], [2]))
print(s.findMedianSortedArrays([1, 2], [3, 4]))
