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

	longestArr, shortestArr := longestArr(nums1, nums2)

	i := 0
	j := 0
	for i < len(longestArr) {
		for j < len(shortestArr) {
			if i == len(longestArr) {
				for ; j < len(shortestArr); j++ {
					finalArr[finalArrIndex] = shortestArr[j]
					finalArrIndex++
				}
				break
			}

			if longestArr[i] < shortestArr[j] {
				finalArr[finalArrIndex] = longestArr[i]
				i++
			} else {
				finalArr[finalArrIndex] = shortestArr[j]
				j++
			}
			finalArrIndex++
		}

		if j == len(shortestArr) {
			for ; i < len(longestArr); i++ {
				finalArr[finalArrIndex] = longestArr[i]
				finalArrIndex++
			}
			break
		}
	}

	if len(finalArr)%2 == 0 {
		return float64((finalArr[medianIndexes[0]] + finalArr[medianIndexes[1]])) / 2.0
	}
	return float64(finalArr[medianIndexes[0]])
}

func longestArr(nums1 []int, nums2 []int) ([]int, []int) {
	if len(nums1) > len(nums2) {
		return nums1, nums2
	}

	return nums2, nums1
}
