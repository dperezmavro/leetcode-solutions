package solution

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	finalArr := make([]int, len(nums1)+len(nums2))
	finalArrIndex := 0
	medianIndexes := make([]int, 2)

	if len(finalArr)%2 == 0 {
		// two indexes
		medianIndexes[0] = (len(finalArr) / 2) - 1
		medianIndexes[1] = len(finalArr) / 2
	} else {
		medianIndexes[0] = len(finalArr) / 2
	}

	i := 0
	j := 0
	for finalArrIndex < len(finalArr) {
		if i == len(nums1) {
			for ; j < len(nums2); j++ {
				finalArr[finalArrIndex] = nums2[j]
				finalArrIndex++
			}
			break
		}
		if j == len(nums2) {
			for ; i < len(nums1); i++ {
				finalArr[finalArrIndex] = nums1[i]
				finalArrIndex++
			}
			break
		}

		if nums1[i] < nums2[j] {
			finalArr[finalArrIndex] = nums1[i]
			i++
		} else {
			finalArr[finalArrIndex] = nums2[j]
			j++
		}
		finalArrIndex++
	}

	if len(finalArr)%2 == 0 {
		return float64((finalArr[medianIndexes[0]] + finalArr[medianIndexes[1]])) / 2.0
	}
	return float64(finalArr[medianIndexes[0]])
}
