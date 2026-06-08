from typing import List


class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        median = 0

        if len(nums1) == 0 and len(nums2) == 0:
            return median

        total_size = len(nums1) + len(nums2)
        result = []
        while len(nums1) != 0 or len(nums2) != 0:
            if len(nums1) == 0:
                result.extend(nums2)
                break

            if len(nums2) == 0:
                result.extend(nums1)
                break

            if nums1[0] < nums2[0]:
                el = nums1.pop(0)
                result.append(el)
            else:
                el = nums2.pop(0)
                result.append(el)

            # TODO(dio): Optimisation - once we are past
            # the midpoint of the result_array, we don't
            # need to carry on - we can just stop there

        if total_size % 2 == 0:
            median = result[total_size // 2] + result[(total_size // 2) - 1]
            median /= 2
        else:
            median = result[total_size // 2]
        return median
