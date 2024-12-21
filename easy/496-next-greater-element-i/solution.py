from typing import List


class Solution:
    def nextGreaterElement(self, nums1: List[int], nums2: List[int]) -> List[int]:
        stack = []
        mappings = dict()

        for idx, n in enumerate(reversed(nums2)):
            if idx == 0:
                mappings[n] = -1
            if stack:
                if n < stack[-1]:
                    mappings[n] = stack[-1]
                else:
                    while stack and n > stack[-1]:
                        stack.pop()
                    if stack:
                        mappings[n] = stack[-1]
                    else:
                        mappings[n] = -1
                
            stack.append(n)
        return [mappings[x] for x in nums1]
