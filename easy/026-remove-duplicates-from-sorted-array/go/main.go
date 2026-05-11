package lib

func removeDuplicates(nums []int) int {
	k := 0
	for i, v := range nums {
		// base case when we start
		if i == 0 {
			k++
			continue
		}

		if nums[k-1] != v {
			nums[k] = v
			k++
		} else {
			nums[k] = -1
		}
	}
	return k
}
