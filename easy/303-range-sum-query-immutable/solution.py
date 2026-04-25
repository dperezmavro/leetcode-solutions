class NumArray:

    def __init__(self, nums: List[int]):
        self.sums = [0] * len(nums)
        self.nums = nums
        for idx, v in enumerate(nums):
            self.sums[idx] = v
            if idx > 0 :
                self.sums[idx] += self.sums[idx-1]

    def sumRange(self, left: int, right: int) -> int:
        return self.sums[right]- self.sums[left] + self.nums[left]
